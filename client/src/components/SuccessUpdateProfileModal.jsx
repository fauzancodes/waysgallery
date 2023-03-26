import { Modal } from "react-bootstrap";

export default function SuccessUpdateProfileModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="success-update-profile-modal" centered>
        <Modal.Body className="!tw-p-7">
          <p className="!tw-text-2xl !tw-text-center !tw-text-custom-success">You have successfully updated your profile.</p>
        </Modal.Body>
      </Modal>
    </>
  );
}