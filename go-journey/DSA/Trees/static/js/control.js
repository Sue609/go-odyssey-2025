/**
 * ⚙️ controls.js
 * 
 * Contsins UI control functions that manage user-triggered actions like 
 * displaying sorted order (In-Order Traversal) and clearing the entire tree.
 * 
 * These are high-;evel utilities that connect user intent to the underlying
 * visualization and API logic.
 */


/**
 * 🧮 showSorted()
 * Fetches the current tree from the backend and performs an **In-Order Traversal** 
 * to show the BST values in sorted order.
 */

async function showSorted() {
    // Fetch the current tree from backend
    const res = await fetch("/tree");
    const tree = await res.json();

    // Helper array, to store sorted values
    const sorted = [];

    // recursive helper function fro in-order traversal
    function inOrder(node) {
        if (!node) return;
        inOrder(node.left);
        sorted.push(node.value);
        inOrder(node.right);
    }


    inorder(tree);


    // Display result on screen
    document.getElementById("sortedDisplay").textContent = "In-Order Traversal(Sorted): " + sorted.join(", ");
}



/**
 * 🧹 clearTree()
 * Displays a modal asking the user for confirmation before clearing the entire tree.
 * 
 * - If confirmed, sends a DELETE request to the backend endpoint `/tree/clear`
 * - On success, it removes all visual elements (nodes and links) from the SVG
 * - If canceled, it simply closes the modal
 */

async function clearTree() {
    // 1️⃣ Create and inject a custom modal overlay into the DOM
    const confrimBox = document.createElement("div");
    confrimBox.innerHTML = `
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


    // 2️⃣ Handle Cancel — remove modal overlay
    document.getElementById("cancelBtn").onclick = () => {
        document.getElementById("confirm-overlay").remove();
    };


    // 3️⃣ Handle Confirm — clear tree from backend and visualization
    document.getElementById("confrimBtn").onclick = async () => {
        try {
            // DELETE backend to clear sorted tree data
            const response = await fetch("/tree/clear", { method: "DELETE "});
            const data = await response.json();
            console.log(data.message);

            // Clear D visualization(remove all SVG elements)
            d3.select("svg").selectAll("*").remove();

            // Optionally, reset UI displays (sorted text, traversal info, etc.)
            document.getElementById("sortedDisplay").textContent = "";

        } catch (err) {
            console.error("Error clearing tree:", err);
        } finally {
            // Remove the confirmation overlay regardless of outcome
            document.getElementById("confirm-overlay").remove();
        }
    };
}