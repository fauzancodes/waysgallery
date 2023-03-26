import { Container, Button } from "react-bootstrap";
import { useNavigate } from "react-router-dom";

function Landing(props) {
  document.title = "WaysGallery | Show Your Work to Inspire Everyone";
  const navigate = useNavigate();
  props.setNavbarOff();

  return (
    <>
      <Container className="tw-min-h-full tw-flex tw-justify-between tw-items-center">
        <div className="tw-flex tw-flex-col tw-w-full md:tw-w-6/12 tw-mx-4 tw-mt-4 md:tw-m-0 tw-z-10">
          <img src="/images/title.webp" alt="WaysGallery" className="tw-mb-4 tw-max-w-xs md:tw-w-7/12"/>
          <h1 className="tw-text-3xl tw-font-bold">Show Your Work to Inspire Everyone</h1>
          <h3 className="tw-text-xl">A place for design creators to gather and share their work with other creators</h3>
          <div className="tw-mt-4">
            {props.isLogin ? (
              <Button onClick={() => navigate("/posts")} className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-mb-4 md:tw-mb-0">Read Posts</Button>
              ) : (
                <>
                  <Button onClick={props.showModalRegister} className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-mb-4 md:tw-mb-0">Join Now</Button>
                  <Button onClick={props.showModalLogin} className="!tw-bg-custom-secondary hover:!tw-bg-custom-secondary-dark !tw-font-bold !tw-text-black !tw-border-0 md:tw-ml-4 tw-w-full md:tw-w-32 tw-mb-4 md:tw-mb-0">Login</Button>
                  <Button variant="link" onClick={() => navigate("/posts")} className="md:tw-ml-4 tw-w-full md:tw-w-48">Just looking around</Button>
                </>
              )
            }
          </div>
        </div>
        <img src="/images/landing-page.webp" alt="landing page" className="animate__animated animate__zoomIn tw-w-6/12 tw-object-contain tw-hidden md:tw-inline" />
        <img src="/images/landing-page-top-left.webp" alt="landing page top left" className="animate__animated animate__slideInDown tw-fixed tw-top-0 tw-left-0 tw-w-2/12 md:tw-w-3/12" />
        <img src="/images/landing-page-bottom-left.webp" alt="landing page bottom left" className="animate__animated animate__slideInLeft tw-fixed tw-bottom-0 tw-left-0 tw-w-2/12 md:tw-w-3/12" />
        <img src="/images/landing-page-bottom-right.webp" alt="landing page bottom-right" className="animate__animated animate__slideInRight tw-fixed tw-bottom-0 tw-right-0 tw-w-2/12 md:tw-w-3/12" />
      </Container>
    </>
  )
}

export default Landing;