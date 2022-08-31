import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";
import EditPlaylistItem from "../components/editPlaylistItem";
import PlaylistItemPreview from "../components/playlistItemPreview";
import { motion } from "framer-motion";
import useSWR from "swr";

const Confirm = () => {
  const [playlist, setPlaylist] = useState(undefined);
  const [editing, setEditing] = useState(false);
  const [editingId, setEditingId] = useState("");
  const [editingIdPlaylistItem, setEditingIdPlaylistItem] = useState(undefined);
  const CreatePlaylist = async () => {};
  useEffect(() => {
    if (playlist) {
      for (let i = 0; i < playlist.length; i++) {
        if (playlist[i].id === editingId) {
          setEditingIdPlaylistItem(playlist[i]);
          setEditing(true);
        }
      }
    }
  }, [editingId]);
  useEffect(() => {
    const generatePlaylist = async () => {
      const params = new URLSearchParams(window.location.search);
      const res = await fetch(
        process.env.NEXT_PUBLIC_BACKEND_URL +
          `/api/generate?playlist=${params.get("playlist")}`,
        {
          credentials: "include",
          mode: "cors",
        }
      ).then((res) => res.json());
      setPlaylist(res);
    };
    generatePlaylist();
  }, []);

  if (playlist === undefined) return "Loading...";
  return (
    <div className="flex flex-col gap-y-8">
      {!editing && <h1 className="font-medium text-6xl">Let's Confirm</h1>}
      {editing && playlist && (
        <EditPlaylistItem
          playlistItem={editingIdPlaylistItem}
          setEditing={setEditing}
        ></EditPlaylistItem>
      )}
      {!editing && playlist && (
        <div className="flex flex-col gap-y-8">
          <div
            className="flex flex-col gap-y-8"
            style={{ maxHeight: "60vh", overflowY: "scroll" }}
          >
            {playlist.map((playlistItem) => {
              return (
                <PlaylistItemPreview
                  playlistItem={playlistItem}
                  setEditing={setEditing}
                  setEditingId={setEditingId}
                  key={playlistItem.id}
                ></PlaylistItemPreview>
              );
            })}
          </div>
          <motion.button
            onClick={CreatePlaylist}
            className="px-16 py-2 border border-black self-end text-2xl"
            whileHover={{
              borderColor: "white",
              backgroundColor: "black",
              color: "white",
            }}
          >
            Create Youtube Playlist
          </motion.button>
        </div>
      )}
    </div>
  );
};

export default Confirm;
