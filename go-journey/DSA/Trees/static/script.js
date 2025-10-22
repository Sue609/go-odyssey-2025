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
  document.getElementById("treeDisplay").textContent = JSON.stringify(tree, null, 2);
}

updateTree();
