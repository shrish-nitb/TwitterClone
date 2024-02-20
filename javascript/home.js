import { tweetClient } from "../client.js";

async function tweet() {
  let content = document.querySelector("#type").value;
  // admin userID here intended to be obtained via cookies
  let user = "18";
  return tweetClient.createTweet({
    tweet: {
      content: content,
      user: user,
    },
  });
}

const form = document.getElementsByTagName("form")[0];
form.onsubmit = (e) => {
  e.preventDefault();
  tweet()
    .then((res) => {
      alert("tweet successful")
    })
    .catch((err) => {
      console.log(err.toString());
    });
};
