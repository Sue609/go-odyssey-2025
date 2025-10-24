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
// a temporary status messade and a fade-in effect for the rendered tree.

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

// Load initial tree on startup
updateTree();

