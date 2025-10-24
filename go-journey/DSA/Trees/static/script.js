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
// Load initial tree on startup
updateTree();

