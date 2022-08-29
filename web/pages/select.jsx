import Head from "next/head";
import Image from "next/image";
import { motion } from "framer-motion";
import SelectPlaylist from "../components/selectplaylist";

const Select = ({ playlists }) => {
  return (
    <div className="flex flex-col gap-y-8">
      <h1 className="font-medium text-6xl">Select your playlist</h1>
      <div className="flex flex-wrap gap-8">
        {playlists.map((playlist) => (
          <SelectPlaylist></SelectPlaylist>
        ))}
      </div>
    </div>
  );
};

export async function getServerSideProps(context) {
  const sessionCookie = context.req.cookies["session_id"];
  console.log(sessionCookie);
  return {
    props: {
      playlists: [],
    },
  };
}
export default Select;
