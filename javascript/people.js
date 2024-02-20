import { tweetClient, userClient } from "../client.js";

async function getUsers() {
  return userClient.listUsers({
    pageSize: 7,
    pageToken: "1",
  });
}

getUsers()
  .then((response) => {
    for (let item of response.users) {
      let userCard = document.createElement("div");
      userCard.classList.add("userCard");
      userCard.innerHTML = `
    <img class="userCardCover" src="https://miro.medium.com/v2/resize:fit:1400/0*oIQlFjlVVkRuXwv6" alt="">
    <img class="userCardPic" src="https://img.freepik.com/premium-photo/smile-happy-proud-business-woman-entrepreneur-leader-working-corporate-office-portrait-accounts-executive-hr-manager-administrator-startup-agency-with-happiness-work_590464-81360.jpg" alt="">
    <div class="userCardInfo">
        <div class="userCardHeading">${item.firstName} ${item.lastName}</div>
        <div class="userCardHeadingRelation">Lorem ipsum dolor sit amet consectetur adipisicing
            elit. Iure deleniti cupiditate error reprehenderit maiores. Praesentium
            exercitationem sunt neque nam sint fugit, dignissimos obcaecati doloremque cumque,
            minima excepturi vero, beatae voluptatibus!</div>
        <div class="userCardActions">

            <div class="userCardBadge"><i class="fa-regular fa-paper-plane" aria-hidden="true"></i>
            </div>
            <div class="userCardBadge"><i class="fa-solid fa-plus" aria-hidden="true"></i></div>
            <div class="userCardBadge"><i class="fa-regular fa-bookmark" aria-hidden="true"></i>
            </div>

        </div>
    </div>
`;
      document.getElementById("collection").appendChild(userCard);
    }
  })
  .catch((error) => {
    alert(error);
  });
