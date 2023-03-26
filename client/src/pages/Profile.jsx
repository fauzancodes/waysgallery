import { Container, Button } from "react-bootstrap";
import { useNavigate, useParams } from "react-router-dom";
import { useState } from "react";
import { API } from '../config/api';

import SuccessFollowingModal from "../components/SuccessFollowingModal";
import AlreadyFollowModal from "../components/AlreadyFollowModal";

export default function Profile(props) {
  props.setNavbarOn();
  const navigate = useNavigate();
  let params = useParams();
  
  const Profile = props.Profiles.find(profile => profile.id === parseInt(params.id));

  document.title = `${props.LoggedInUserId === parseInt(params.id) ? "Profie" : Profile.name} | WaysGallery`;

  const handleFollow = async (target_id, current_user_id) => {
    if (props.isLogin === true) {
      try {
        const response = await API.post('/following/' + target_id);
        const newFollowData = {
          id: response.data.data.id,
          follwoing_id: target_id,
          user: {id:current_user_id},
        }
        props.SetFollowings([...props.Followings, newFollowData]);
        setModalSuccessFollowingShow(true);
      }
      catch (error) {
        if (error.response.data.message === "Sorry, you are already following this user")  setModalAlreadyFollowShow(true);
      }
    }
    else props.showModalLogin();
  }
  const handleHire = (target_id) => {
    if (props.isLogin === true) {
      navigate(`/order/${target_id}`);
    }
    else props.showModalLogin();
  }

  const [modalSuccessFollowingShow, setModalSuccessFollowingShow] = useState(false);
  const [modalAlreadyFollowShow, setModalAlreadyFollowShow] = useState(false);

  return (
    <>
      <SuccessFollowingModal 
        show={modalSuccessFollowingShow} 
        onHide={() => setModalSuccessFollowingShow(false)} 
      />
      <AlreadyFollowModal  
        show={modalAlreadyFollowShow} 
        onHide={() => setModalAlreadyFollowShow(false)} 
      />
      <Container className="tw-min-h-full">
        <img src="/images/icon-profile-background.png" alt="Latest Post Background" className="tw-absolute tw-right-0 tw-top-0 tw-z-[-1] tw-hidden lg:tw-block" />
        <div className="tw-mt-28 tw-flex tw-flex-wrap tw-justify-between tw-items-center">
          <div className="animate__animated animate__slideInLeft tw-w-full md:tw-w-5/12 tw-flex tw-flex-col tw-items-center md:tw-items-start">
            {Profile?.image === "" || Profile?.image === undefined || Profile?.image === null ? (
              <img src="/images/icon-user.png" alt="Profile Icon" className="!tw-rounded-full !tw-cursor-pointer !tw-object-cover !tw-w-28 !tw-h-28 !tw-bg-custom-secondary tw-mb-4"/>
            ) : (
              <img src={Profile.image} alt="Profile Icon" className="!tw-rounded-full !tw-cursor-pointer !tw-object-cover !tw-w-28 !tw-h-28 !tw-bg-custom-secondary tw-mb-4"/>
            )}
            <h3 className="tw-font-bold">{Profile.name}</h3>
            {Profile.greeting !== "" && (
              <h1 className="tw-font-bold">{Profile.greeting}</h1>
            )}
            <div className="tw-flex tw-flex-wrap tw-my-12">
              {props.LoggedInUserId === parseInt(params.id) ? (
                <Button onClick={() => navigate("/update-profile")} className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-mr-4 tw-mb-4">Edit Profile</Button>
              ) : (
                <>
                  <Button onClick={() => handleFollow(parseInt(params.id), props.LoggedInUserId)} className="!tw-bg-custom-secondary hover:!tw-bg-custom-secondary-dark !tw-font-bold !tw-text-black !tw-border-0 tw-w-full md:tw-w-32 tw-mr-4 tw-mb-4">Follow</Button>
                  <Button onClick={() => handleHire(parseInt(params.id))} className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-mr-4 tw-mb-4">Hire</Button>
                </>
              )}
            </div>
          </div>
          {
            props.Posts.filter(post => post.user.id === parseInt(params.id))[0] !== undefined ? (
              <img src={props.Posts.filter(post => post.user.id === parseInt(params.id))[0].image_1} alt="Latest Post" className="animate__animated animate__zoomIn tw-w-full md:tw-w-6/12" />
            ) : (
              <img src="/images/icon-profile-background.png" alt="Latest Post" className="tw-w-full md:tw-w-6/12" />
            )
          }
        </div>
        <h3 className="tw-mt-20 tw-mb-10 tw-font-bold">My Works</h3>
        <div className="tw-flex tw-flex-wrap tw-mb-28">
          {
            Profile.arts.length > 0 ? (
              <>
                {Profile.arts.sort((a, b) => b.id - a.id).map((art, index) => (
                  <img key={index} onClick={() => props.handleShowModalImageEnlarger(art.image)} src={art.image} alt={`Work ${index+1}`} className="animate__animated animate__zoomIn tw-object-cover tw-p-2.5 tw-rounded tw-w-full md:tw-w-3/12 tw-h-60 tw-cursor-pointer" />
                ))}
                {
                  props.LoggedInUserId === parseInt(params.id) && (
                    <div onClick={() => navigate("/upload-art")} className="tw-p-2.5 tw-rounded tw-w-full md:tw-w-3/12 tw-h-60 tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-border-dashed tw-flex tw-justify-center tw-items-center tw-cursor-pointer">
                      <img src="/images/icon-cloud-upload.png" alt="Add More" className="tw-object-contain tw-w-6/12"/>
                    </div>
                  )
                }
              </>
            ) : (
              props.LoggedInUserId === parseInt(params.id) ? (
                <div onClick={() => navigate("/upload-art")} className="tw-p-2.5 tw-rounded tw-w-full md:tw-w-3/12 tw-h-60 tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-border-dashed tw-flex tw-justify-center tw-items-center tw-cursor-pointer">
                  <img src="/images/icon-cloud-upload.png" alt="Add More" className="tw-object-contain tw-w-6/12"/>
                </div>
              ) : (
                <p className="tw-opacity-50">There are no works to display.</p>
              )
            )
          }
        </div>
      </Container>
    </>
  );
}