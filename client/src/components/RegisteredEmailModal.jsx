import { Modal } from "react-bootstrap";

export default function RegisteredEmailModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="registered-email-modal" centered>
        <Modal.Body className="!tw-p-7">
          <p className="!tw-text-2xl !tw-text-center !tw-text-custom-danger">This email is already registered, please login.</p>
        </Modal.Body>
      </Modal>
    </>
  );
}