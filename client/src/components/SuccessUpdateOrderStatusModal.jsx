import { Modal } from "react-bootstrap";

export default function SuccessUpdateOrderStatusModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="success-update-profile-modal" centered>
        <Modal.Body className="!tw-p-7">
          <p className="!tw-text-2xl !tw-text-center !tw-text-custom-success">{props.message}</p>
        </Modal.Body>
      </Modal>
    </>
  );
}