import { Modal } from "react-bootstrap";

export default function SuccessRegisterModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="success-update-product-modal" centered>
        <Modal.Body className="!tw-p-7">
          <p className="!tw-text-2xl !tw-text-center !tw-text-custom-success">Your account is successfully registered.</p>
        </Modal.Body>
      </Modal>
    </>
  );
}