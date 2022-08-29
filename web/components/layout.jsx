import React from "react";

const Layout = ({ children }) => {
  return (
    <div className="w-screen lg:px-32 lg:py-16">
      <div className="w-full flex justify-center flex-col gap-y-24">
        <h3 className="font-medium text-3xl">Concertify</h3>
        <div className="lg:w-[1400px]">{children}</div>
      </div>
    </div>
  );
};

export default Layout;
