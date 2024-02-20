import * as connectrpcconnect from 'https://esm.run/@connectrpc/connect';
import * as connectrpcconnectWeb from 'https://esm.run/@connectrpc/connect-web';
import { UserService } from './internal/gen/api/users_connect.js';
import { TweetService } from './internal/gen/api/tweets_connect.js';

const transport = connectrpcconnectWeb.createConnectTransport({
  baseUrl: "http://127.0.0.1:8080",
  useGet: true,
});

export const userClient = await connectrpcconnect.createPromiseClient(UserService, transport);
export const tweetClient = await connectrpcconnect.createPromiseClient(TweetService, transport);
