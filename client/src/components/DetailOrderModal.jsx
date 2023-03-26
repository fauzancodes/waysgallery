import { Modal } from "react-bootstrap";

export default function DetailOrderModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="success-upload-post-modal" centered>
        <Modal.Body className="!tw-p-7">
          <p className="!tw-text-2xl">Title: {props.title}.</p>
          <p className="!tw-text-2xl">Description: {props.description}.</p>
          <p className="!tw-text-2xl !tw-font-bold !tw-text-custom-success">Price: {props.price.toLocaleString("id-ID", {style: "currency",currency: "IDR"})}</p>
        </Modal.Body>
      </Modal>
    </>
  );
}