import { Modal } from "react-bootstrap";

export default function ImageEnlargerModal(props) {
  return (
    <>
      <Modal {...props} size="lg" aria-labelledby="success-upload-post-modal" centered>
        <Modal.Body className="!tw-flex !tw-flex-col !tw-justify-center !tw-items-center">
          <img src={props.imageToEnlarge} alt="Preview Enlarged" className="tw-rounded" />
          <a href={props.imageToEnlarge} target="_blank" rel="noreferrer" className="tw-no-underline tw-bg-custom-success hover:tw-bg-custom-success-dark !tw-text-white tw-mt-4 tw-rounded tw-px-4 tw-py-2">Download</a>
        </Modal.Body>
      </Modal>
    </>
  );
}