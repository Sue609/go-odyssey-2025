/**
 * Handles all tree generation and traversal-related logic.
 * - Builds a BTS from user input (comma-separated values).
 * - Calls backend routes to insert values and retrieve updates tree data.
 * - Animates the process with fade-in transition and status messages.
 * - Includes traversal visualization for inorder, preorder and postorder.
 * 
 */


// === Tree Generation ===
// Builds a new BST based on user input and animates the process.
async function generateTree() {
    const input = document.getElementById("numbersInput").value;

    if (!input) {
        showStatus("Please enter some numbers first.", "error");
        return;
    }

    // Parse input into numbers
    const numbers = input
        .spilit(",")
        .map(n => parseInt(n.trim()))
        .filter(n => !isNaN(n));

    if (numbers.length === 0) {
        showStatus("Please enter valid numbers separated by commas.", "error");
        return;
    }

    // Show loading animation while generating
    showStatus("🌳 Generating your beautiful tree...", "loading");
    
    // Clear any existing tree data on backend
    await fetch('/tree/clear', { method: 'DELETE' });


    // Insert each number sequentially for smooth visual feedback
    for (const num of numbers) {
        await fetch(`/insert?value=${num}`, { method: 'POST' });
        await new Promise(resolve => setTimeout(resolve, 200)); // tiny delay for animation
    }

    // Fetch updated tree and visualize
    const response = await fetch('/tree');
    const treeData = await response.json();

    // Clear any previous visualization and apply fade-in animation
    d3.select('svg').selectAll('*').remove();
    d3.select('svg')
        .style("opacity", 0)
        .transition()
        .duration(800)
        .style("opacity", 1);

    
    renderTree(treeData);
    showStatus("🌸 Tree generated successfully!", "success");
    console.log("Generated tree with:", numbers); 
    
}


// === Traversal Visualization ===
// Allows users to select traversal type and visualize the order dynamically.
document.getElementById("traverseBtn").addEventListener("click", async () => {
  const type = document.getElementById("traversalType").value;
  const resultBox = document.getElementById("traversalResult");

  if (!type) {
    resultBox.textContent = "Please select a traversal type!";
    resultBox.style.color = "red";
    return;
  }

  try {
    const res = await fetch(`/traverse?type=${type}`);
    if (!res.ok) throw new Error("Traversal failed");

    const data = await res.json();
    const nodes = data.result;

    resultBox.textContent = `${type.toUpperCase()} Traversal: ${nodes.join(" → ")}`;
    resultBox.style.color = "#333";

    highlightTraversal(nodes);
  } catch (err) {
    resultBox.textContent = "Error fetching traversal.";
    resultBox.style.color = "red";
    console.error(err);
  }
});


