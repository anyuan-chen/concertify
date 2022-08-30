import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import { motion } from "framer-motion";
import { useRouter } from "next/router";

const Home: NextPage = () => {
  const Container = {
    rest: {
      opacity: 0,
    },
    load: {
      opacity: 1,
      transition: {
        staggerChildren: 0.5,
      },
    },
  };
  const Element = {
    rest: {
      opacity: 0,
    },
    load: {
      opacity: 1,
    },
  };
  const router = useRouter();
  const login = () => {
    router.push(process.env.NEXT_PUBLIC_BACKEND_URL + "/spotify/login");
  };
  return (
    <motion.div
      initial="rest"
      animate="load"
      variants={Container}
      className="flex flex-col gap-y-16"
    >
      <div className="flex flex-col gap-y-4">
        <div className="flex gap-x-2">
          <motion.h2 className="text-3xl" variants={Element}>
            Feeling broke?
          </motion.h2>{" "}
          <motion.h2 className="text-3xl" variants={Element}>
            Feeling sad?
          </motion.h2>
        </div>
        <motion.h1 className="text-6xl font-medium" variants={Element}>
          Turn your Spotify playlist into a list of live performances.
        </motion.h1>
      </div>
      <div className="flex gap-x-4">
        <motion.img src="/images/concert1.png" variants={Element}></motion.img>
        <motion.img src="/images/concert2.png" variants={Element}></motion.img>
        <motion.img src="/images/concert3.png" variants={Element}></motion.img>
      </div>
      <button
        className="self-start text-4xl px-8 py-4 font-medium border border-black"
        onClick={login}
      >
        Login With Spotify
      </button>
    </motion.div>
  );
};

export default Home;
