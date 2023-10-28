import Pusher from "pusher-js";

const pusherKey = process.env.NEXT_PUBLIC_PUSHER_APP_KEY as string;
const pusherCluster = process.env.NEXT_PUBLIC_PUSHER_APP_CLUSTER as string;

const pusher = new Pusher(pusherKey, {
  cluster: pusherCluster,
});

export default pusher;
