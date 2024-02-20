import { tweetClient, userClient } from "../client.js";

async function deleteAccount() {
  return userClient.deleteUser({
    id: document.getElementById("delUserId").value,
  });
}

async function updateAccount() {
  return userClient.updateUser({
    id: document.getElementById("updUserId").value,
    firstName: document.querySelector("#fn").value,
    lastName: document.querySelector("#ln").value,
  });
}

document.getElementById("deletionForm").onsubmit = (e) => {
  e.preventDefault();
  deleteAccount()
    .then((response) => {
      alert("deletion successful");
    })
    .catch((error) => {
      alert(error);
    });
};

document.getElementById("updationForm").onsubmit = (e) => {
  e.preventDefault();
  updateAccount()
    .then((response) => {
        alert(`changed to ${response.firstName} ${response.lastName}`)
    })
    .catch((error) => {
      alert(error);
    });
};
