import React, { useState } from 'react';
import { Calculate } from "./components/Calculate";

function App() {
  const [text, setText] = useState<string>("");

  return (
    <div style={{
      textAlign: "center",
    }}>
      <Calculate />
    </div>
  );
}

export default App;
