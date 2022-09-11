import React from "react";
import { useState } from "react";
import { useEffect } from "react";
const Success = () => {
  const [link, setLink] = useState("");
  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    setLink(`youtube.com/playlist?list=${params.get("playlist")}`);
  }, []);

  return (
    <div className="flex flex-col gap-y-8">
      <h1 className="text-6xl font-medium">Here's your playlist</h1>
      <div className="flex border border-black py-4 px-4 items-center justify-between">
        <a
          href={"https://www." + link}
          target="_blank"
          className="text-5xl font-medium underline"
        >
          {link}
        </a>
        <button
          onClick={() => {
            navigator.clipboard.writeText("https://www." + link);
          }}
        >
          <img src="./images/clipboard-copy.svg"></img>
        </button>
      </div>
      <h2
        className="text-3xl"
        style={{ position: "absolute", bottom: "150px" }}
      >
        Feel free to share this site. Check out{" "}
        <a
          href="https://andrewchen.tech"
          style={{ textDecoration: "underline" }}
        >
          andrewchen.tech
        </a>{" "}
        for my other work.
      </h2>
    </div>
  );
};

export default Success;
