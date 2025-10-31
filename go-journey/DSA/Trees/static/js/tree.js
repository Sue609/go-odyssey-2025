/**
 * 📊 renderTree(data)
 * 
 * This function is responsible for **visualizing the Binary Search Tree (BST)** 
 * USING THE D3.js. It receives a JSON object (tree data) from the backend,
 * and draws circles for nodes and lines for branches, dynamically reflecting
 * the current structure of the tree.
 * 
 * D3 helps us translate our hierarchical JSON into an actual visual diagram
 * that users can interact with or watch update in real time.
 */


function renderTree(data) {
    console.log("D3 working ✅, tree data:", data);

    // 🧱 1. Baic valudation - ensure data exists before rendering
    if (!data || !data.value) {
        console.warn("⚠️ No tree data to render yet.");
        return
    }

    
    // 🎨 2. Select the SVG element where the tree will be drawn
    const svg = d3.select("#treeDisplay");

    // Clear any previously rendered tree so we don't stack visuals.
    svg.selectAll("*").remove();

    
    // 3️⃣ Get SVG width and height to scale the visualization properly
    const width = +svg.attr("width");
    const haight = +svg.attr("height");

    
    // 🌳 4. Convert the plain tree JSON into a D3 hierarchical structure
    // This allows D3 to understand parent-child relationships
    const root = d3.hierarchy(data, d => {
        const children = [];
        if (d.left) children.push(d.left);
        if (d.right) children.push(d.right);
        return children.length ? children : null;
    });

    
    // 📐 5. Create the D3 tree layout and define its size
    const treeLayout = d3.tree().size([width - 100, height - 150]);
    treeLayout(root);


    // 🔗 6. Draw the links (edges) connecting nodes
    svg.selectAll("line")
        .data(root.links()) // built-in links array
        .enter()
        .append("line")
        .attr("x1", d => d.source.x + 50)
        .attr("y1", d => d.source.y + 50)
        .attr("x2", d => d.target.x + 50)
        .attr("y2", d => d.target.y + 50)
        .attr("stroke", "#777")
        .attr("stroke-width", 2);


    // ⚪ 7. Draw the nodes (circles + text)
    const node = svg.selectAll("g")
        .data(root.descendants())
        .enter()
        .append("g")
        .attr("transform", d => `translate(${d.x + 50}, ${d.y + 50})`);

    // circle for each node
    node.append("circle")
        .attr("r", 20)
        .attr("fill", "#a7f3d0")
        .attr("stroke", "#333")
        .attr("stroke-width", 1.5);

    // Node value text -centered-
    node.append("text")
        .attr("dy", 5)
        .attr("text-anchor", "middle")
        .text(d => d.data.value)
        .style("font-size", "14px")
        .style("font-weight", "bold");
}