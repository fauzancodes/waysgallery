import { Modal } from "react-bootstrap";

export default function SuccessFollowingModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="success-follow-modal" centered>
        <Modal.Body className="!tw-p-7">
          <p className="!tw-text-2xl !tw-text-center !tw-text-custom-success">You have successfully followed this account.</p>
        </Modal.Body>
      </Modal>
    </>
  );
}