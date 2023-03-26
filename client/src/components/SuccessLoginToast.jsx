import { Toast, Alert } from "react-bootstrap";

export default function SuccessLoginToast(props) {
  return (
    <Toast onClose={props.onClose} show={props.show} delay={5000} autohide className="!tw-fixed !tw-bottom-8 !tw-left-4 !tw-z-50 !tw-w-fit">
      <Alert key={"success"} variant={"success"} className="!tw-text-lg !tw-m-0 !tw-w-fit !tw-bg-custom-primary !tw-text-black">
        Login Success!
      </Alert>
    </Toast> 
  )
}