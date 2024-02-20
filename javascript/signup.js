import { userClient } from "../client.js";

async function register() {
  let fn = document.querySelector("#fn").value;
  let ln = document.querySelector("#ln").value;
  return userClient.createUser({
    user: {
      firstName: fn,
      lastName: ln,
    },
  });
}

const form = document.getElementsByTagName("form")[0];
form.onsubmit = (e) => {
  e.preventDefault();
  register()
    .then((res) => {
      alert(JSON.stringify(res));
    })
    .catch((err) => {
      alert(err.toString());
    });
};
