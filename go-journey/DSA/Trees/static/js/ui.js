/**
 * ui.js
 *
 * Contains helper functions for updating user interface feedback.
 * Handles success, error, and loading states with fade-in/fade-out animations.
 */

// Displays temporary status messages (loading, success, error)
function showStatus(message, type) {
  const status = document.getElementById("statusMessage");
  status.textContent = message;

  if (type === "loading") {
    status.style.color = "#2e8b57"; // greenish
  } else if (type === "success") {
    status.style.color = "#28a745"; // bright green
  } else if (type === "error") {
    status.style.color = "#d9534f"; // red
  }

  status.style.opacity = 1;
  setTimeout(() => {
    status.style.transition = "opacity 0.5s ease";
    status.style.opacity = 0;
  }, 2000);
}
