import { tweetClient, userClient } from "../client.js";

async function getUserName() {}

async function list() {
  return tweetClient.listTweets({
    pageSize: 7,
    pageToken: "1",
    // concerned userID here intended to be obtained via url params 
    user: "18",
  });
}

list()
  .then(async (res) => {
    for (const i of res.tweets) {
      let user = await userClient.getUser({
        id: i.user,
      });

      var cardDiv = document.createElement("div");
      cardDiv.className = "card";

      cardDiv.innerHTML = `
      <div>
          <div class="cardHeader">
              <img class="avatar" src="https://w0.peakpx.com/wallpaper/837/397/HD-wallpaper-johnny-bravo-thumbnail.jpg" alt="" />
              <div class="usrBadge" first>
                  <div>
                      ${user.firstName} ${user.lastName}
                      <div>@${user.firstName} ${user.lastName}</div>
                      <div>1d ago</div>
                  </div>
              </div>
              <div class="postMenu">
                  <i class="fa-solid fa-ellipsis-vertical"></i>
              </div>
          </div>
          <div class="tweetContent">
             ${i.content}
          </div>
      </div>
      <div class="interactTray">
          <div>
              <i class="fa-regular fa-comment"></i>
              <div>61</div>
          </div>
          <div>
              <i class="fa-solid fa-retweet"></i>
              <div>12</div>
          </div>
          <div loved>
              <i class="fa-solid fa-heart"></i>
              <div>6.2k</div>
          </div>
          <div>
              <i class="fa-solid fa-share-nodes"></i>
              <div>61</div>
          </div>
      </div>
  `;

      document.getElementById("feed").appendChild(cardDiv);
    }
  })
  .catch((err) => {
    alert(err.toString());
  });
