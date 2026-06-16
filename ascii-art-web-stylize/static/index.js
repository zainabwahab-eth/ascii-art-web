const clearBtn = document.querySelector(".clear-row");
const result = document.querySelector(".result");
const text = document.getElementById("text");

clearBtn.addEventListener("click", () => {
  result.textContent = "";
  text.textContent = "";
});
