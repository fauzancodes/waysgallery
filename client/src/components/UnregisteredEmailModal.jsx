import { Modal } from "react-bootstrap";

export default function UnregisteredEmailModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="success-add-product-modal" centered>
        <Modal.Body className="!tw-p-7">
          <p className="!tw-text-2xl !tw-text-center !tw-text-custom-danger">Sorry, the email you entered was not registered, please register first.</p>
        </Modal.Body>
      </Modal>
    </>
  );
}