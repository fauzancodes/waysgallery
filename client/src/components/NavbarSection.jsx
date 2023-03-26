import { React } from "react";
import { Button, Container, Navbar, Dropdown } from 'react-bootstrap';
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

export default function NavbarSection(props) {
  const navigate = useNavigate();

  const Profile = props.Profiles.find(profile => profile.id === props.LoggedInUserId);

  return (
    <>
      <Navbar collapseOnSelect expand="md" fixed="top" bg="light" variant="light" className="tw-border-b-2 tw-border-custom-secondary-dark">
        <Container>
          <Link to="/">
            <Navbar.Brand className="tw-flex tw-items-center">
              <img src="/images/title.webp" alt="WaysGallery" className="tw-h-16"/>
            </Navbar.Brand>
          </Link>
          <Navbar.Toggle aria-controls="responsive-navbar-nav" />
          <Navbar.Collapse id="responsive-navbar-nav" className="tw-flex tw-justify-between md:tw-justify-end">
            {
              props.isLogin === true ? (
                <>
                  <Button onClick={() => navigate("/upload-post")} className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 !tw-w-32 !tw-my-4 md:!tw-my-0 md:!tw-mr-4 !tw-order-2 md:tw-order-1">Upload</Button>
                  <Dropdown className="!tw-w-fit !tw-order-1 md:!tw-order-2">
                    <Dropdown.Toggle id="profile-menu" className="!tw-border-0 !tw-bg-transparent">
                      {Profile?.image === undefined || Profile?.image === ""? (
                        <img src="/images/icon-user.png" alt="Profile Icon" className="!tw-rounded-full !tw-cursor-pointer !tw-object-cover !tw-w-14 !tw-h-14 !tw-bg-custom-secondary"/>
                      ) : (
                        <img src={Profile.image} alt="Profile Icon" className="!tw-rounded-full !tw-cursor-pointer !tw-object-cover !tw-w-14 !tw-h-14 !tw-bg-custom-secondary"/>
                      )}
                    </Dropdown.Toggle>
                    <Dropdown.Menu className="!tw-border-0 !tw-shadow-[0_0.25rem_0.25rem_rgba(0,0,0,0.25),0.25rem_0.25rem_1.25rem_rgba(0,0,0,0.25)]">
                      <Dropdown.Item onClick={() => navigate(`/profile/${props.LoggedInUserId}`)} className="!tw-font-bold !tw-flex !tw-items-center !tw-py-2">
                        <img src="/images/icon-user.png" alt="Profile" className="!tw-mr-2 !tw-w-6"/>
                        Profile
                      </Dropdown.Item>
                      <Dropdown.Item onClick={() => navigate("/order-list")} className="!tw-font-bold !tw-flex !tw-items-center !tw-py-2">
                        <img src="/images/icon-order.png" alt="Order" className="!tw-mr-2 !tw-w-6"/>
                        Order
                      </Dropdown.Item>
                      <Dropdown.Divider/>
                      <Dropdown.Item onClick={props.logout} className="!tw-font-bold !tw-flex !tw-items-center !tw-py-2">
                        <img src="/images/icon-logout.png" alt="Profile" className="!tw-mr-2 !tw-w-6"/>
                        Logout
                      </Dropdown.Item>
                    </Dropdown.Menu>
                  </Dropdown>
                </>
              ) : (
                <>
                  <Button onClick={props.showModalRegister} className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-my-4 md:tw-my-0">Join Now</Button>
                  <Button onClick={props.showModalLogin} className="!tw-bg-custom-secondary hover:!tw-bg-custom-secondary-dark !tw-font-bold !tw-text-black !tw-border-0 md:tw-ml-4 tw-w-full md:tw-w-32 tw-mb-4 md:tw-mb-0">Login</Button>
                </>
              )
            }
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </>
  );
}