import React from "react";
import Form from "./components/form";
function App() {
  return (
    <div className="App">
      <form
        encType="multipart/form-data"
        action="http://localhost:8080/upload"
        method="post"
      >
        <input type="file" name="myFile" />
        <input type="submit" value="upload" />
      </form>
      <Form />
    </div>
  );
}

export default App;
