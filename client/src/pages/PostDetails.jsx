import { Container, Button } from "react-bootstrap";
import { useNavigate, useParams, Link } from "react-router-dom";
import { useState } from "react";
import { API } from '../config/api';

import SuccessFollowingModal from "../components/SuccessFollowingModal";
import AlreadyFollowModal from "../components/AlreadyFollowModal";

function PostDetails(props) {
  const navigate = useNavigate();
  props.setNavbarOn();

  const params = useParams();
  const Posts = props.Posts;
  let Post = Posts.find(post => post.id === parseInt(params.id));
  const Profiles = props.Profiles;
  let Profile = Profiles.find(profile => profile.id === Post.user.id);

  document.title = `${Post.title} | WaysGallery`;
  
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
        <div className="tw-mt-28 tw-flex tw-items-center tw-justify-between">
          <Link to={`/profile/${Post.user.id}`} className="!tw-no-underline">
            <div className="tw-flex tw-items-center tw-mb-8">
              {
                Profile.image === "" ? (
                  <img src="/images/icon-user.png" alt="Profile Icon" className="!tw-rounded-full !tw-cursor-pointer !tw-object-cover !tw-w-14 !tw-h-14 !tw-bg-custom-secondary"/>
                ) : (
                  <img src={Profile.image} alt="Profile Icon" className="!tw-rounded-full !tw-cursor-pointer !tw-object-cover !tw-w-14 !tw-h-14 !tw-bg-custom-secondary"/>
                )
              }
              <div className="tw-ml-4">
                <div className="tw-font-bold">{Post.title}</div>
                <div>{Post.user.name}</div>
              </div>
            </div>
          </Link>
          {
            Post.user.id !== props.User.id && (
              <div>
                <Button Button onClick={() => handleFollow(Post.user.id, props.User.id)} className="!tw-bg-custom-secondary hover:!tw-bg-custom-secondary-dark !tw-font-bold !tw-text-black !tw-border-0 tw-w-full md:tw-w-32 md:tw-mb-0">Follow</Button>
                <Button onClick={() => handleHire(Post.user.id)} className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-my-4 md:tw-my-0 md:tw-ml-4">Hire</Button>
              </div>
            )
          }
        </div>
        {
          Post.image_1 === "" ? (
            <img src="/images/image-placeholder" alt="Thumbnail" className="tw-w-full tw-mb-4 tw-h-60 md:tw-h-[30rem] tw-object-cover tw-rounded" />
          ) : (
            <img onClick={() => props.handleShowModalImageEnlarger(Post.image_1)} src={Post.image_1} alt="Thumbnail" className="animate__animated animate__zoomIn tw-w-full tw-mb-4 tw-h-60 md:tw-h-[30rem] tw-object-cover tw-rounded tw-cursor-pointer" />
          )
        }
        <div className="tw-flex tw-mb-8">
          {
            Post.image_2 !== "" && (
              <img onClick={() => props.handleShowModalImageEnlarger(Post.image_2)} src={Post.image_2} alt="Thumbnail" className="animate__animated animate__zoomIn tw-w-3/12 tw-h-20 md:tw-h-40 tw-object-cover tw-p-1 tw-rounded tw-cursor-pointer" />
            )
          }
          {
            Post.image_3 !== "" && (
              <img onClick={() => props.handleShowModalImageEnlarger(Post.image_3)} src={Post.image_3} alt="Thumbnail" className="animate__animated animate__zoomIn tw-w-3/12 tw-h-20 md:tw-h-40 tw-object-cover tw-p-1 tw-rounded tw-cursor-pointer" />
            )
          }
          {
            Post.image_4 !== "" && (
              <img onClick={() => props.handleShowModalImageEnlarger(Post.image_4)} src={Post.image_4} alt="Thumbnail" className="animate__animated animate__zoomIn tw-w-3/12 tw-h-20 md:tw-h-40 tw-object-cover tw-p-1 tw-rounded tw-cursor-pointer" />
            )
          }
          {
            Post.image_5 !== "" && (
              <img onClick={() => props.handleShowModalImageEnlarger(Post.image_5)} src={Post.image_5} alt="Thumbnail" className="animate__animated animate__zoomIn tw-w-3/12 tw-h-20 md:tw-h-40 tw-object-cover tw-p-1 tw-rounded tw-cursor-pointer" />
            )
          }
        </div>
        <h5 className="animate__animated animate__zoomIn tw-font-bold tw-mb-4">👋 Say Hello <a className="tw-no-underline" href={`mailto:${Post.user.email}`}>{Post.user.email}</a></h5>
        <p>{Post.description}</p>
      </Container>
    </>
  )
}

export default PostDetails;