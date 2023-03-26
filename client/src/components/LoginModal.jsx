import { Modal, Form, Button } from "react-bootstrap";

export default function LoginModal(props) {
  return (
    <>
      <Modal {...props} aria-labelledby="login-modal" centered>
        <Modal.Body>
          <h1 id="login-modal" className="!tw-font-bold !tw-text-custom-primary !tw-px-4 !tw-py-6">
            Login
          </h1>
          <Form onSubmit={props.loginOnSubmit} className="!tw-px-4 !tw-pb-4">
            <Form.Group className="!tw-mb-6" controlId="formEmail">
              <Form.Control type="email" onChange={props.loginOnChange} name="email" value={props.formLogin.email} placeholder="Email" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
            </Form.Group>
            <Form.Group className="!tw-mb-6" controlId="formPassword">
              <Form.Control type="password" onChange={props.loginOnChange} name="password" value={props.formLogin.password} placeholder="Password" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
            </Form.Group>
            <Button type="submit" className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-w-full !tw-font-bold !tw-text-lg tw-mb-2 !tw-mt-2 !tw-p-4 !tw-border-0">Login</Button>
            <Form.Text className="!tw-text-black !tw-text-lg !tw-text-center !tw-block">
              Don't have an account ? Klik <strong onClick={props.changeModal} className="!tw-cursor-pointer hover:tw-text-custom-primary hover:tw-underline">Here</strong>
            </Form.Text>
          </Form>
        </Modal.Body>
      </Modal>
    </>
  );
}