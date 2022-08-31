import React from "react";
import { motion } from "framer-motion";

const YoutubeVideoPreview = ({ src, title, artist, href }) => {
  return (
    <div className="flex gap-x-8">
      <img style={{ height: "100px", width: "150px" }} src={src}></img>
      <div className="flex flex-col justify-between">
        <motion.a
          href={href ? href : ""}
          className="flex gap-x-8"
          target="_blank"
        >
          <motion.h2
            className="text-3xl font-medium"
            whileHover={{ textDecoration: "underline" }}
          >
            {title.replace(/&quot;/g, '"')}
          </motion.h2>
        </motion.a>

        <h3 className="text-xl font-medium text-gray-600">{artist}</h3>
      </div>
    </div>
  );
};

export default YoutubeVideoPreview;
