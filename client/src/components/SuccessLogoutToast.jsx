import { Toast, Alert } from "react-bootstrap";

export default function SuccessLogoutToast(props) {
  return (
    <Toast onClose={props.onClose} show={props.show} delay={5000} autohide className="!tw-fixed !tw-top-4 !tw-right-4 !z-50 !tw-w-fit">
      <Alert key={"secondary"} variant={"secondary"} className="!tw-text-lg !tw-m-0 !tw-w-fit !tw-bg-custom-secondary !tw-text-black">
        Logout Success!
      </Alert>
    </Toast> 
  )
}