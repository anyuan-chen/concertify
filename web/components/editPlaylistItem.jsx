import React from "react";
import YoutubeVideoPreview from "./youtubeVideoPreview";
import Checkbox from "../components/checkbox";
import { motion } from "framer-motion";

const EditPlaylistItem = ({
  playlistItem,
  setEditing,
  editingId,
  setPlaylistItem,
  playlist,
  setPlaylist,
}) => {
  const [checkedId, setCheckedId] = React.useState(
    playlistItem.youtube_search_response[0].id.videoId
  );
  const [customUrl, setCustomUrl] = React.useState("");
  const SaveAndExit = async () => {
    let newPlaylist;
    if (customUrl !== "") {
      const res = await fetch(
        process.env.NEXT_PUBLIC_BACKEND_URL + `/api/video?query=${customUrl}`,
        {
          credentials: "include",
          mode: "cors",
        }
      );
      const data = await res.json();
      data.id = {
        videoId: data.id,
      };
      newPlaylist = [data, ...playlistItem.youtube_search_response];
      newPlaylist = newPlaylist.slice(0, newPlaylist.length - 1);
    } else {
      let editingIndex = 0;
      for (let i = 0; i < playlistItem.youtube_search_response.length; i++) {
        if (playlistItem.youtube_search_response[i].id.videoId === checkedId) {
          editingIndex = i;
        }
      }
      newPlaylist = [...playlistItem.youtube_search_response];
      const tempPlaylistItem =
        playlistItem.youtube_search_response[editingIndex];
      newPlaylist[editingIndex] = playlistItem.youtube_search_response[0];
      newPlaylist[0] = tempPlaylistItem;
    }
    setPlaylist(
      playlist.map((item) => {
        if (item.id === playlistItem.id) {
          return {
            ...item,
            youtube_search_response: newPlaylist,
          };
        }
        return item;
      })
    );
    setEditing(false);
  };
  if (playlistItem) {
    console.log(playlistItem.name);
  }
  return (
    playlistItem && (
      <div className="flex flex-col p-8 border border-black gap-y-8">
        <div className="flex gap-x-4 items-end">
          <h3 className="text-3xl font-medium">
            {playlistItem.name.replace(/&quot;/g, '\\"')}
          </h3>
          <h3 className="text-xl text-grey-600">
            {playlistItem.artists
              .reduce((prev, cur) => {
                return [...prev, cur.name];
              }, "")
              .join(", ")}
          </h3>
        </div>
        <form className="flex flex-col gap-y-8">
          {playlistItem.youtube_search_response.map((response, index) => {
            return (
              <div
                className="flex justify-between items-center"
                key={response.id.videoId}
              >
                <YoutubeVideoPreview
                  href={`https://www.youtube.com/watch?v=${response.id.videoId}`}
                  src={response.snippet.thumbnails.default.url}
                  title={response.snippet.title}
                  artist={response.snippet.channelTitle}
                ></YoutubeVideoPreview>
                <Checkbox
                  checked={
                    checkedId ===
                    playlistItem.youtube_search_response[index].id.videoId
                      ? true
                      : false
                  }
                  setChecked={() => {
                    setCustomUrl("");
                    setCheckedId(
                      playlistItem.youtube_search_response[index].id.videoId
                    );
                  }}
                ></Checkbox>
              </div>
            );
          })}
        </form>
        <div className="flex justify-between pt-8">
          <div className="flex flex-col gap-y-2">
            <label className="text-xl font-medium">
              Have a specific performance in mind?
            </label>
            <input
              type="text"
              placeholder="youtube link"
              className="w-[500px]"
              value={customUrl}
              onChange={(e) => {
                setCustomUrl(e.target.value);
                setCheckedId(e.target.value);
              }}
            ></input>
          </div>
          <motion.button
            onClick={SaveAndExit}
            className="px-16 py-2 border border-black self-end text-2xl"
            whileHover={{
              borderColor: "white",
              backgroundColor: "black",
              color: "white",
            }}
          >
            Save
          </motion.button>
        </div>
      </div>
    )
  );
};

export default EditPlaylistItem;
