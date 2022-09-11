import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";
import EditPlaylistItem from "../components/editPlaylistItem";
import PlaylistItemPreview from "../components/playlistItemPreview";
import { motion } from "framer-motion";
import useSWR from "swr";

const Confirm = () => {
  const router = useRouter();
  const [playlist, setPlaylist] = useState(undefined);
  const [editing, setEditing] = useState(false);
  const [editingId, setEditingId] = useState("");
  const [editingIdPlaylistItem, setEditingIdPlaylistItem] = useState(undefined);
  const [editingTrigger, setEditingTrigger] = useState(false);
  const CreatePlaylist = async () => {
    const res = await fetch(
      process.env.NEXT_PUBLIC_BACKEND_URL + `/api/create`,
      {
        method: "POST",
        credentials: "include",
        mode: "cors",
        body: JSON.stringify(playlist),
      }
    );
    const data = await res.json();
    console.log(res.status);
    if (res.status === 200) {
      router.push(`/success?playlist=${data.id}`);
    }
  };
  useEffect(() => {
    if (playlist) {
      for (let i = 0; i < playlist.length; i++) {
        if (playlist[i].id === editingId) {
          setEditingIdPlaylistItem(playlist[i]);
          setEditing(true);
        }
      }
    }
  }, [editingId, editingTrigger]);
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
          editingId={editingId}
          playlist={playlist}
          setPlaylist={setPlaylist}
        ></EditPlaylistItem>
      )}
      {!editing && playlist && (
        <div className="flex flex-col gap-y-8">
          <div
            className="flex flex-col gap-y-8"
            style={{ maxHeight: "60vh", overflowY: "scroll" }}
          >
            {playlist.map((playlistItem, index) => {
              return (
                <PlaylistItemPreview
                  playlistItem={playlistItem}
                  setEditing={setEditing}
                  setEditingId={setEditingId}
                  setEditingTrigger={setEditingTrigger}
                  editingTrigger={editingTrigger}
                  key={index}
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
