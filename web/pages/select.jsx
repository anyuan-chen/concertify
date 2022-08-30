import Head from "next/head";
import Image from "next/image";
import { motion } from "framer-motion";
import useSWR from "swr";
import SelectPlaylist from "../components/selectplaylist";
import { useState } from "react";
import { useRouter } from "next/router";

const Select = () => {
  const router = useRouter();
  const fetcher = (url, args) => fetch(url, args).then((res) => res.json());
  const [page, setPage] = useState(1);
  const [submitting, setSubmitting] = useState(false);
  const { data, error } = useSWR(
    [
      process.env.NEXT_PUBLIC_BACKEND_URL + `/api/playlists?page=${page}`,
      {
        credentials: "include",
        mode: "cors",
      },
    ],
    fetcher
  );
  const SubmitSelect = async (id) => {
    setSubmitting(true);
    router.push(`/confirm?playlist=${id}`);
  };
  if (error) return "An error has occurred.";
  if (!data) return "Loading...";
  return (
    <div className="flex flex-col gap-y-8">
      <div className="flex">
        <h1 className="font-medium text-6xl">Select your playlist</h1>
      </div>
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
      <div className="flex justify-between">
        <button
          onClick={() => setPage(page - 1)}
          disabled={page === 0}
          className={`px-8 py-2 border  ${
            page === 0
              ? "text-gray-500 border-gray-500"
              : "text-black border-black"
          }`}
        >
          Prev Page
        </button>
        <button
          onClick={() => setPage(page + 1)}
          className={`px-8 py-2 border border-black ${
            Math.ceil(data.total / 20) - 1 === page
              ? "text-gray-500 border-gray-500"
              : "text-black border-black"
          }`}
          disabled={Math.ceil(data.total / 20) - 1 === page}
        >
          Next Page
        </button>
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
