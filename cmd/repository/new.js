import React from "react";
import Form from "./components/Form";

class UploadModal extends React.Component {
  constructor(props) {
    super(props);
    //add a ref value to your state and a setter to set the ref
    this.setDropZoneRef = this.setDropZoneRef.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  onSubmit() {
    // call open file select dialog if haven't select any file
    //here use the formref you've set
    this.formRef.submit();
  }
  handleSubmit(values) {
    //handling submit
  }
  render() {
    return (
      <div>
        <p>Upload files</p>
        <Form ref={ref => (this.formRef = ref)} onSubmit={this.handleSubmit} />
        <Button onClick={this.onSubmit}>Upload</Button>
      </div>
    );
  }
}
