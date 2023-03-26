import { Modal } from "react-bootstrap";

export default function FailedMakeOrderModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="success-add-product-modal" centered>
        <Modal.Body className="!tw-p-7">
          <p className="!tw-text-2xl !tw-text-center !tw-text-custom-danger">Something went wrong, please try again.</p>
        </Modal.Body>
      </Modal>
    </>
  );
}