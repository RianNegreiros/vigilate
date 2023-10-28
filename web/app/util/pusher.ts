import Pusher from 'pusher-js';

const pusherKey = process.env.NEXT_PUBLIC_PUSHER_KEY as string;
const pusherCluster = process.env.NEXT_PUBLIC_PUSHER_CLUSTER as string;

const pusher = new Pusher("573dec410087e0ec6559", {
  cluster: "sa1",
});

export default pusher;
