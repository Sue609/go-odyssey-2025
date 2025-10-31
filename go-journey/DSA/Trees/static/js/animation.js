/**
 * animation.js
 * 
 * Handles all visual animations for the Binary Search Tree (BST):
 * - Highlights traversal paths using D3 transitions.
 * - Uses Canvas API for lightweight node drawing and traversal simulation.
 * - Adds interactive buttons for pausing, resuming, and generating random trees.
 * 
 * Dependencies:
 * - D3.js (for tree visualization)
 * - Functions from render.js (renderTree)
 * - Functions from api.js (updateTree)
 */


// === D3 Traversal Highlight ===
// Smoothly highlights nodes in the chosen traversal order.
function highlightTraversal(nodes) {
  if (!nodes || nodes.length === 0) return;

  const circles = d3.selectAll("circle");

  // Reset all nodes to their default color
  circles.attr("fill", "#a7f3d0");

  // Sequentially highlight each node in traversal order
  nodes.forEach((value, i) => {
    setTimeout(() => {
      circles
        .filter(d => d.data && d.data.value === value)
        .transition()
        .duration(400)
        .attr("fill", "#ffb703") // Highlight color
        .transition()
        .duration(400)
        .attr("fill", "#a7f3d0"); // Return to normal
    }, i * 800); // Delay between highlights
  });
}



// === Canvas Setup ===
const canvas = document.getElementById("treeCanvas");
const ctx = canvas.getContext("2d");

let animationPaused = false;
let traversalData = [];
let currentIndex = 0;


// === Node Drawing ===
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



// === Tree Drawing ===
function drawTree(sequence) {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  const gap = canvas.width / (sequence.length + 1);

  sequence.forEach((val, i) => {
    drawNode((i + 1) * gap, canvas.height / 2, val);
  });
}


// === Traversal Animation ===
async function animateTraversal(sequence) {
  for (let i = 0; i < sequence.length; i++) {
    while (animationPaused) await new Promise(r => setTimeout(r, 100));

    drawTree(sequence.slice(0, i));
    drawNode(
      (i + 1) * (canvas.width / (sequence.length + 1)),
      canvas.height / 2,
      sequence[i],
      true
    );

    await new Promise(r => setTimeout(r, 800));
  }
}



// === Traversal Buttons ===
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

    drawTree([]);
    animateTraversal(traversalData);

  } catch (error) {
    alert("Error fetching traversal data.");
    console.error(error);
  }
});



// === Pause / Resume Buttons ===
document.getElementById("pauseBtn").addEventListener("click", () => {
  animationPaused = true;
});
document.getElementById("resumeBtn").addEventListener("click", () => {
  animationPaused = false;
});



// === Generate Random Tree ===
// Quickly creates a random BST for testing traversal visualizations.
document.getElementById("generateRandomBtn").addEventListener("click", async () => {
  // Generate 5–10 random numbers between 1–100
  const randomCount = Math.floor(Math.random() * 6) + 5;
  const randomNumbers = Array.from({ length: randomCount }, () =>
    Math.floor(Math.random() * 100)
  );

  console.log("Generated numbers:", randomNumbers);

  // Clear existing tree data
  await fetch("/tree/clear", { method: "DELETE" });

  // Sequentially insert numbers with small delay
  for (const num of randomNumbers) {
    await fetch(`/insert?value=${num}`, { method: "POST" });
    await new Promise(resolve => setTimeout(resolve, 200));
  }

  // Fetch updated tree and visualize
  const response = await fetch("/tree");
  const data = await response.json();

  d3.select("svg").selectAll("*").remove();
  renderTree(data);

  alert(`🌳 New Random Tree Generated:\n[${randomNumbers.join(", ")}]`);
});


// === Load initial tree when page opens ===
updateTree();