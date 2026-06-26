const clearBtn = document.querySelector(".clear-row");
const result = document.querySelector(".result");
const text = document.getElementById("text");
const downloadBtn = document.querySelector(".download-btn");

clearBtn.addEventListener("click", () => {
  result.textContent = "";
  text.textContent = "";
  downloadBtn.classList.add("hidden");
  clearBtn.classList.add("hidden");
});
