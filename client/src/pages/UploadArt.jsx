import { Container, Form, Button } from "react-bootstrap";
import { useState } from "react";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom"
import { API } from '../config/api';

export default function UploadPost(props) {
  document.title = "Upload Post | WaysGallery";
  props.setNavbarOn();
  const navigate = useNavigate();

  const [image, setImage] = useState("");
  const [formArt, setFormArt] = useState({
    image: "",
  });
  const formArtHandleOnChange = (e) => {
    setFormArt({
      ...formArt,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage(url);
    }
  };
  const formArtHandleOnSubmit = useMutation(async (e) => {
      e.preventDefault();
  
      const config = {
        headers: {
          "Content-type": "multipart/form-data",
        },
      };

      const formData = new FormData();
      formData.set('image', formArt.image[0], formArt.image[0].name);
      try {
        const response = await API.post('/art', formData, config);

        const newArt = {
          id: response.data.data.id,
          image: response.data.data.image,
        }
        props.SetProfiles(prevState => prevState.map(profile => {
          if (profile.id === props.LoggedInUserId) {
            return {
              ...profile,
              arts: [...profile.arts, newArt]
            };
          }
          return profile;
        }));
  
        setFormArt((formArt) => ({
          ...formArt,
          image: "",
        }));
        setImage("");
        props.showModalSuccessAddArt();
        navigate(`/profile/${props.LoggedInUserId}`);
      }
      catch (error) {
        return
      }
  });

  return (
    <Container className="tw-min-h-full">
      <Form onSubmit={(e) => formArtHandleOnSubmit.mutate(e)} className="animate__animated animate__zoomIn md:tw-flex md:tw-flex-col md:tw-items-center tw-mt-28 tw-mb-20">
        {image === "" ? (
          <label for="image" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-full md:tw-w-6/12 tw-cursor-pointer">
            <img src="/images/icon-cloud-upload.png" alt="Preview" className="tw-object-contain"/>
            <p><strong className="tw-text-custom-primary">Browse</strong> to choose a file</p>
          </label>
        ) : (
          <label for="image" className="tw-w-full md:tw-w-6/12 tw-flex tw-justify-center tw-cursor-pointer">
            <img src={image} alt="Preview" className="tw-object-contain tw-w-full"/>
          </label>
        )}
        <input id="image" onChange={(e) => formArtHandleOnChange(e)} name="image" type="file" className="tw-hidden"/>
        <Button type="submit" className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32 tw-mt-10">Save</Button>
      </Form>
    </Container>
  );
}