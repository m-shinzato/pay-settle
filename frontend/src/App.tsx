import React, { useState } from 'react';
import { TrackButton } from "./components/TrackButton";
import { TrackText } from "./components/TrackText";
import { Calculate } from "./components/Calculate";

function App() {
  const [text, setText] = useState<string>("");
  const onClick = () => {
    setText("hello track");
  }

  return (
    <div style={{
      textAlign: "center",
    }}>
      <Calculate />
      <TrackButton onClick={onClick} />
      <TrackText text={text}/>
    </div>
  );
}

export default App;
