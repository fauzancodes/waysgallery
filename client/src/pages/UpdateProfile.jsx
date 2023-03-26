import { Container, Form, Button } from "react-bootstrap";
import { useState } from "react";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import { API } from '../config/api';

export default function UpdateProfile(props) {
  document.title = "Update Profile | WaysGallery";
  props.setNavbarOn();
  const navigate = useNavigate();

  const Profile = props.Profiles.find(profile => profile.id === props.LoggedInUserId);

  const [image, setImage] = useState(Profile?.image);
  const [formProfile, setFormProfile] = useState({
    name: props.User.name,
    greeting: Profile?.greeting,
    image: image,
  });
  const formProfileHandleOnChange = (e) => {
    setFormProfile({
      ...formProfile,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage(url);
    }
  };
  const formProfileHandleOnSubmit = useMutation(async (e) => {
      e.preventDefault();
  
      const config = {
        headers: {
          "Content-type": "multipart/form-data",
        },
      };

      const formData = new FormData();
      if (formProfile.image) {
        formData.set('image', formProfile?.image[0], formProfile?.image[0]?.name);
      }
      formData.set('name', formProfile.name);
      formData.set('greeting', formProfile.greeting);

      try {
        const response = await API.patch('/profile', formData, config);

        const newProfile = {
          name: response.data.data.name,
          greeting: response.data.data.greeting,
          image: response.data.data.image,
        }
        props.SetProfiles(prevState => prevState.map(profile => {
          if (profile.id === props.LoggedInUserId) {
            return {
              ...profile,
              name: newProfile.name,
              greeting: newProfile.greeting,
              image: newProfile.image,
            };
          }
          return profile;
        }));

        setFormProfile((formProfile) => ({
          ...formProfile,
          name: Profile.name,
          greeting: Profile.greeting,
          image: image,
        }));
        setImage("");
        props.showModalSuccessUpdateProfile(true);
        navigate(`/profile/${props.LoggedInUserId}`);
      }
      catch (error) {
        return
      }
  });

  return (
    <Container className="tw-min-h-full tw-flex tw-justify-between tw-pt-28">
      {
        image === undefined || image === null || image === "" ? (
          <div className="animate__animated animate__slideInLeft tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-5 tw-border-dashed tw-rounded tw-hidden md:tw-flex tw-justify-center tw-items-center tw-w-6/12">
            <img src="/images/icon-camera.png" alt="Profile" className="tw-object-contain"/>
          </div>
        ) : (
          <div className="animate__animated animate__slideInLeft tw-rounded tw-hidden md:tw-flex tw-justify-center tw-items-center tw-w-6/12 tw-h-96 tw-overflow-hidden">
            <img src={image} alt="Profile" className="tw-w-full tw-object-cover"/>
          </div>
        )
      }
      <Form onSubmit={(e) => formProfileHandleOnSubmit.mutate(e)} className="animate__animated animate__slideInRight !tw-px-4 !tw-pb-4 tw-flex tw-flex-col tw-items-center tw-w-full md:tw-w-fit">
        {
          image === undefined || image === null || image === "" ? (
            <label for="profile-picture" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-5 tw-border-dashed tw-rounded-full tw-mb-6 tw-cursor-pointer">
              <img src="/images/icon-camera.png" alt="Profile" className="tw-object-contain"/>
            </label>
          ) : (
            <label for="profile-picture" className="tw-rounded-full tw-mb-6 tw-cursor-pointer">
              <img src={image} alt="Profile" className="tw-rounded-full tw-w-40 tw-h-40 tw-object-cover"/>
            </label>
          )
        }
        <input id="profile-picture" type="file" onChange={(e) => formProfileHandleOnChange(e)} name="image" className="tw-hidden" />
        <Form.Group className="!tw-mb-6 tw-w-full" controlId="formGreeting">
          <Form.Control type="text" onChange={(e) => formProfileHandleOnChange(e)} name="greeting" value={formProfile.greeting} placeholder="Greeting" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
        </Form.Group>
        <Form.Group className="!tw-mb-6 tw-w-full" controlId="formPassword">
          <Form.Control type="text" onChange={(e) => formProfileHandleOnChange(e)} name="name" value={formProfile.name} placeholder="Full Name" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" required/>
        </Form.Group>
        <Button type="submit" className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32">Save</Button>
      </Form>
    </Container>
  );
}