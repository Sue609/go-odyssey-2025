/*
    This file handles API call (insert, delete, traverse, etc.)
    It acts a bridege between our front end and backend API.

*/

/**
 * Handles inserting a new node into the BST.
 * 
 * 1️⃣ Reads user's input value from the text field.
 * 2️⃣ Sends a GET request to the backend API endpoint `/insert?value=<userInput>`.
 * 3️⃣ Once insertion is done, it calls '`updateTree()` to fetch and redraw the updated tree.
 * 
 * This function represents the 'Create' operation in our tree's CRUD system
 */
async function insertValue() {
    const val = document.getElementById("valueInput").value;
    await fetch(`/insert?value=${val}`); // Sends insert request to API
    await updateTree();
}



/**
 * Handles deleting an existing node from the BST.
 * 
 * 1️⃣ Reads the user's input value from the text field.
 * 2️⃣ Sends a get request to the backend endpoint `/delete?value=<userInput>`
 * 3️⃣ Once the node is deleted, it calls 'updateTree() to re-render the latest tree state.
 * 
 * This represents the "Delete" operation in our tree's CRUD logic.
 */

async function deleteValue() {
    const val = document.getElementById("valueInput").value;
    await fetch(`/delete?value=${val}`);
    await updateTree();

}



/**
 * Fetches the current tree structure from the backend and triggers a re-render
 * 
 * 1️⃣ Sends a GET request to `/tree` to retrieve the latest tree in JSON Format.
 * 2️⃣ Passes the fetched to 'renderTree(tree)', which draw it on the screen.
 * 
 * This acts as the "sychronization" function between the backend and frontend visualization.
 */
async function updateTree() {
    const res = await fetch("/tree");
    const tree = await res.json(); // convert response to JSON
    renderTree(tree);
}