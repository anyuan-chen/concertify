import React from "react";

const Checkbox = ({ checked, setChecked }) => {
  return (
    <input
      type="checkbox"
      checked={checked}
      onChange={() => setChecked(!checked)}
      className="text-black"
      style={{
        height: "50px",
        width: "50px",
        backgroundImage: "url(/images/checkbox.png)",
      }}
    />
  );
};

export default Checkbox;
