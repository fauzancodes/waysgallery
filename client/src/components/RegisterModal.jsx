import { Modal, Form, Button } from "react-bootstrap";

export default function RegisterModal(props) {
  return (
    <Modal {...props} aria-labelledby="register-modal" centered>
      <Modal.Body>
        <h1 id="register-modal" className="!tw-font-bold !tw-text-custom-primary !tw-px-4 !tw-py-6">
          Register
        </h1>
        <Form onSubmit={props.registerOnSubmit} className="!tw-px-4 !tw-pb-4">
          <Form.Group className="!tw-mb-6" controlId="formEmail">
            <Form.Control type="email" onChange={props.registerOnChange} name="email" value={props.formRegister.email} placeholder="Email" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
          </Form.Group>
          <Form.Group className="!tw-mb-6" controlId="formPassword">
            <Form.Control type="password" onChange={props.registerOnChange} name="password" value={props.formRegister.password} placeholder="Password" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
          </Form.Group>
          <Form.Group className="!tw-mb-6" controlId="formText">
            <Form.Control type="text" onChange={props.registerOnChange} name="name" value={props.formRegister.name} placeholder="Full Name" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
          </Form.Group>
          <Button type="submit" className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-w-full !tw-font-bold !tw-text-lg tw-mb-2 !tw-mt-2 !tw-p-4 !tw-border-0">Register</Button>
          <Form.Text className="!tw-text-black !tw-text-lg !tw-text-center !tw-block">
            Already have an account ? Klik <strong onClick={props.changeModal} className="!tw-cursor-pointer hover:tw-text-custom-primary hover:tw-underline">Here</strong>
          </Form.Text>
        </Form>
      </Modal.Body>
    </Modal>
  );
}