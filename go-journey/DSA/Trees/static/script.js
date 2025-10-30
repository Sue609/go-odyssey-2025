async function insertValue() {
  const val = document.getElementById("valueInput").value;
  await fetch(`/insert?value=${val}`);
  await updateTree();
}

async function deleteValue() {
  const val = document.getElementById("valueInput").value;
  await fetch(`/delete?value=${val}`);
  await updateTree();
}

async function updateTree() {
  const res = await fetch("/tree");
  const tree = await res.json();
  renderTree(tree);
}

function renderTree(data) {
  console.log("D3 working ✅, tree data:", data);

  if (!data || !data.value) {
    console.warn("⚠️ No tree data to render yet.");
    return;
  }

  const svg = d3.select("#treeDisplay");
  svg.selectAll("*").remove(); // clear previous render

  const width = +svg.attr("width");
  const height = +svg.attr("height");

  // Convert our Go BST shape into D3 hierarchy format
  const root = d3.hierarchy(data, d => {
    const children = [];
    if (d.left) children.push(d.left);
    if (d.right) children.push(d.right);
    return children.length ? children : null;
  });

  // D3 tree layout
  const treeLayout = d3.tree().size([width - 100, height - 150]);
  treeLayout(root);

  // Draw links (edges)
  svg.selectAll("line")
    .data(root.links())
    .enter()
    .append("line")
    .attr("x1", d => d.source.x + 50)
    .attr("y1", d => d.source.y + 50)
    .attr("x2", d => d.target.x + 50)
    .attr("y2", d => d.target.y + 50)
    .attr("stroke", "#777")
    .attr("stroke-width", 2);

  // Draw nodes (circles + text)
  const node = svg.selectAll("g")
    .data(root.descendants())
    .enter()
    .append("g")
    .attr("transform", d => `translate(${d.x + 50}, ${d.y + 50})`);

  node.append("circle")
    .attr("r", 20)
    .attr("fill", "#a7f3d0")
    .attr("stroke", "#333")
    .attr("stroke-width", 1.5);

  node.append("text")
    .attr("dy", 5)
    .attr("text-anchor", "middle")
    .text(d => d.data.value)
    .style("font-size", "14px")
    .style("font-weight", "bold");
}

async function showSorted() {
  const res = await fetch("/tree");
  const tree = await res.json();

  const sorted = [];
  function inOrder(node) {
    if (!node) return;
    inOrder(node.left);
    sorted.push(node.value);
    inOrder(node.right)
  }

  inOrder(tree);

  document.getElementById("sortedDisplay").textContent = "In-Order Traversal (Sorted): " + sorted.join(", ");
}


// Function for clearing the entire tree
async function clearTree() {
  // Create confirmation modal
  const confirmBox = document.createElement("div");
  confirmBox.innerHTML = `
    <div id="confirm-overlay" style="
      position: fixed; inset: 0; background: rgba(0,0,0,0.4);
      display: flex; align-items: center; justify-content: center;
      z-index: 9999;
    ">
      <div style="
        background: white; padding: 20px 30px; border-radius: 12px;
        text-align: center; box-shadow: 0 4px 20px rgba(0,0,0,0.2);
        max-width: 320px;
      ">
        <h3 style="margin-bottom: 10px;">Are you sure?</h3>
        <p style="font-size: 14px; color: #555;">This will permanently clear your tree data.</p>
        <div style="margin-top: 20px; display: flex; gap: 10px; justify-content: center;">
          <button id="cancelBtn" style="padding: 8px 16px; border: none; border-radius: 8px; background: #e5e5e5;">Cancel</button>
          <button id="confirmBtn" style="padding: 8px 16px; border: none; border-radius: 8px; background: #ef4444; color: white;">Yes, Clear</button>
        </div>
      </div>
    </div>
  `;
  document.body.appendChild(confirmBox);

  // Handle buttons
  document.getElementById("cancelBtn").onclick = () => {
    document.getElementById("confirm-overlay").remove();
  };

  document.getElementById("confirmBtn").onclick = async () => {
    try {
      const response = await fetch("/tree/clear", { method: "DELETE" });
      const data = await response.json();
      console.log(data.message);

      // Clear the D3 visualization
      d3.select("svg").selectAll("*").remove();
    } catch (err) {
      console.error("Error clearing tree:", err);
    } finally {
      document.getElementById("confirm-overlay").remove();
    }
  };
}


// Function that creates a mini animation while building
// a temporary status message and a fade-in effect for the rendered tree.

async function generateTree() {
  const input = document.getElementById("numbersInput").value;
  const status = document.getElementById("statusMessage");

  if (!input) {
    showStatus("Please enter some numbers first!", "error");
    return;
  }

  // Split input and convert to integers
  const numbers = input
    .split(",")
    .map(n => parseInt(n.trim()))
    .filter(n => !isNaN(n));

  if (numbers.length === 0) {
    showStatus("Please enter valid numbers separated by commas.", "error");
    return;
  }

  // Show loading animation
  showStatus("🌳 Generating your beautiful tree...", "loading");

  // Clear current tree before generating a new one
  await fetch('/tree/clear', { method: 'DELETE' });

  // Insert each number one by one (with tiny delay for fun animation)
  for (const num of numbers) {
    await fetch(`/insert?value=${num}`, { method: 'POST' });
    await new Promise(resolve => setTimeout(resolve, 200)); // slow down slightly
  }

  // Fetch updated tree data
  const response = await fetch('/tree');
  const treeData = await response.json();

  // Clear old visualization
  d3.select('svg').selectAll('*').remove();

  // Fade-in animation for new tree
  d3.select('svg')
    .style("opacity", 0)
    .transition()
    .duration(800)
    .style("opacity", 1);

  renderTree(treeData);

  showStatus("🌸 Tree generated successfully!", "success");
  console.log("Generated tree with:", numbers);
}

// Helper to show messages with color and animation
function showStatus(message, type) {
  const status = document.getElementById("statusMessage");
  status.textContent = message;

  if (type === "loading") {
    status.style.color = "#2e8b57";
  } else if (type === "success") {
    status.style.color = "#28a745";
  } else if (type === "error") {
    status.style.color = "#d9534f";
  }

  status.style.opacity = 1;
  setTimeout(() => {
    status.style.transition = "opacity 0.5s ease";
    status.style.opacity = 0;
  }, 2000);
}



// === Traversal Visualization ===
document.getElementById("traverseBtn").addEventListener("click", async () => {
  const type = document.getElementById("traversalType").value;
  const resultBox = document.getElementById("traversalResult");

  if (!type) {
    resultBox.textContent = "Please select a traversal type!";
    resultBox.style.color = "red";
    return;
  }

  // Fetch traversal order from Go backend
  try {
    const res = await fetch(`/traverse?type=${type}`);
    if  (!res.ok) throw new Error("Traversal failed");
    const data = await res.json();

    const nodes = data.result;
    resultBox.textContent= `${type.toUpperCase()} Traversal: ${nodes.join(" → ")}`;
    resultBox.style.color = "#333";

    // Animate traversal on tree
    highlightTraversal(nodes);

  } catch (err) {
    resultBox.textContent = "Error fetching traversal."
    resultBox.style.color = "red"
    console.error(err);
  }
});


// Highlight traversal animation
function highlightTraversal(nodes) {
  if (!nodes || nodes.length === 0) return;

  let circles = d3.selectAll("circle");

  // Reset all circles to default color first
  circles.attr("fill", "#4CAF50");

  nodes.forEach((value, i) => {
    setTimeout(() => {
      circles
        .filter(d => d.data && d.data.value === value)
        .transition()
        .duration(400)
        .attr("fill", "#ffb703") // Highlight color
        .transition()
        .duration(400)
        .attr("fill", "#a7f3d0"); // Back to default
    }, i * 800); // time gap between node highlights
  });
}




const canvas = document.getElementById("treeCanvas");
const ctx = canvas.getContext("2d");

let animationPaused = false;
let traversalData = [];
let currentIndex = 0;

// Draw a simple node
function drawNode(x, y, value, highlight = false) {
    ctx.beginPath();
    ctx.arc(x, y, 25, 0, Math.PI * 2);
    ctx.fillStyle = highlight ? "#90EE90" : "#2196F3";
    ctx.fill();
    ctx.strokeStyle = "#333";
    ctx.stroke();
    ctx.fillStyle = "white";
    ctx.font = "16px sans-serif";
    ctx.textAlign = "center";
    ctx.textBaseline = "middle";
    ctx.fillText(value, x, y);
}

// Clear and redraw tree layout
function drawTree(sequence) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    const gap = canvas.width / (sequence.length + 1);
    sequence.forEach((val, i) => {
        drawNode((i + 1) * gap, canvas.height / 2, val);
    });
}

// Animate traversal
async function animateTraversal(sequence) {
    for (let i = 0; i < sequence.length; i++) {
        while (animationPaused) await new Promise(r => setTimeout(r, 100));
        drawTree(sequence.slice(0, i));
        drawNode((i + 1) * (canvas.width / (sequence.length + 1)), canvas.height / 2, sequence[i], true);
        await new Promise(r => setTimeout(r, 800));
    }
}

document.getElementById("traverseBtn").addEventListener("click", async () => {
    const traversalType = document.getElementById("traversalType").value;

    if (!traversalType) {
        alert("Please select a traversal type first!");
        return;
    }

    try {
        const response = await fetch(`/traverse?type=${traversalType}`);
        const data = await response.json();

        traversalData = data.result;
        currentIndex = 0;
        animationPaused = false;

        drawTree([]); // Clear old
        animateTraversal(traversalData);

    } catch (error) {
        alert("Error fetching traversal data.");
        console.error(error);
    }
});

// Pause/Resume logic
document.getElementById("pauseBtn").addEventListener("click", () => {
    animationPaused = true;
});
document.getElementById("resumeBtn").addEventListener("click", () => {
    animationPaused = false;
});



document.getElementById("generateRandomBtn").addEventListener("click", async () => {
    // Generate between 5 and 10 random numbers between 1–100
    const randomCount = Math.floor(Math.random() * 6) + 5;
    const randomNumbers = Array.from({ length: randomCount }, () =>
        Math.floor(Math.random() * 100)
    );

    console.log("Generated numbers:", randomNumbers);

    // ✅ Clear the existing tree using your actual endpoint
    await fetch("/tree/clear", { method: "DELETE" });

    // Sequentially insert each random number (with small delay for fun)
    for (const num of randomNumbers) {
        await fetch(`/insert?value=${num}`, { method: "POST" });
        await new Promise(resolve => setTimeout(resolve, 200)); // animation delay
    }

    // ✅ Fetch updated tree data
    const response = await fetch("/tree");
    const data = await response.json();

    // ✅ Call your renderTree() to visualize it
    d3.select("svg").selectAll("*").remove(); // clear current
    renderTree(data);

    alert(`🌳 New Random Tree Generated:\n[${randomNumbers.join(", ")}]`);
});

// Load initial tree on startup
updateTree();
