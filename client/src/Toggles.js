import React, { useState, useEffect, useRef } from "react";
import { TreeSelect } from "primereact/treeselect";
import { NodeService } from "./service/NodeService";
import "primereact/resources/themes/lara-light-cyan/theme.css";
import "primereact/resources/primereact.min.css";
import "primeicons/primeicons.css";
import Upload from "./FileUpload";

export default function Preview() {
  const [nodes, setNodes] = useState(null);
  const [selectedNodeKey, setSelectedNodeKey] = useState(null);

  useEffect(() => {
    NodeService.getTreeNodes().then((data) => setNodes(data));
  }, []);

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "flex-start",
        height: "100vh",
        paddingTop: "10%",
      }}
    >
      <TreeSelect
        value={selectedNodeKey}
        onChange={(e) => setSelectedNodeKey(e.value)}
        options={nodes}
        className="md:w-20rem w-full"
        placeholder="Select Filter"
      ></TreeSelect>
      <Upload selectedFilter={selectedNodeKey} />
    </div>
  );
}
