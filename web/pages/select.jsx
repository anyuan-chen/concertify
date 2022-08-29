import Head from "next/head";
import Image from "next/image";
import { motion } from "framer-motion";
import useSWR from "swr";
import SelectPlaylist from "../components/selectplaylist";

const Select = () => {
  const fetcher = (url, args) => fetch(url, args).then((res) => res.json());
  const { data, error } = useSWR(
    [
      process.env.NEXT_PUBLIC_BACKEND_URL + "/api/playlists",
      {
        credentials: "include",
        mode: "cors",
      },
    ],
    fetcher
  );
  if (error) return "An error has occurred.";
  if (!data) return "Loading...";
  return (
    <div className="flex flex-col gap-y-8">
      <h1 className="font-medium text-6xl">Select your playlist</h1>
      <div className="flex flex-wrap gap-8">
        {data?.items?.map((playlist) => (
          <button>
            <SelectPlaylist
              src={playlist?.images[0]?.url}
              title={playlist?.name}
              author={playlist?.owner?.display_name}
            ></SelectPlaylist>
          </button>
        ))}
      </div>
    </div>
  );
};

// export async function getServerSideProps(context) {
//   const sessionCookie = context.req.cookies["session_id"];
// const req = await fetch(
//   `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/playlists`,
//   {
//     credentials: "include",
//     mode: "cors",
//     method: "GET",
//   }
// );
//   console.log(req);
//   //const data = await req.json();
//   //console.log(data);
//   return {
//     props: {
//       playlists: [],
//     },
//   };
// }
export default Select;
