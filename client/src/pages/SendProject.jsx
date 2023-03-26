import { Container, Form, Button } from "react-bootstrap";
import { useState } from "react";
import { useMutation } from "react-query";
import { useNavigate, useParams } from "react-router-dom";
import { API } from '../config/api';

export default function SendProject(props) {
  document.title = "Send Project | WaysGallery";
  props.setNavbarOn();
  const navigate = useNavigate();
  let params = useParams();

  const [image1, setImage1] = useState("");
  const [image2, setImage2] = useState("");
  const [image3, setImage3] = useState("");
  const [image4, setImage4] = useState("");
  const [image5, setImage5] = useState("");
  const [formProject, setFormProject] = useState({
    description: "",
    image1: "",
    image2: "",
    image3: "",
    image4: "",
    image5: "",
  });
  const formProjectHandleOnChange = (e) => {
    setFormProject({
      ...formProject,
      [e.target.name]: e.target.value,
    });
  };
  const formProjectHandleOnChangeImage1 = (e) => {
    setFormProject({
      ...formProject,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage1(url);
    }
  };
  const formProjectHandleOnChangeImage2 = (e) => {
    setFormProject({
      ...formProject,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage2(url);
    }
  };
  const formProjectHandleOnChangeImage3 = (e) => {
    setFormProject({
      ...formProject,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage3(url);
    }
  };
  const formProjectHandleOnChangeImage4 = (e) => {
    setFormProject({
      ...formProject,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage4(url);
    }
  };
  const formProjectHandleOnChangeImage5 = (e) => {
    setFormProject({
      ...formProject,
      [e.target.name]:
        e.target.type === 'file' ? e.target.files : e.target.value,
    });
    if (e.target.type === 'file') {
      let url = URL.createObjectURL(e.target.files[0]);
      setImage5(url);
    }
  };
  const formProjectHandleOnSubmit = useMutation(async (e) => {
      e.preventDefault();
  
      const config = {
        headers: {
          "Content-type": "multipart/form-data",
        },
      };

      const formData = new FormData();
      if (formProject.image1) {
        formData.set('image1', formProject?.image1[0], formProject?.image1[0]?.name);
      }
      if (formProject.image2) {
        formData.set('image2', formProject?.image2[0], formProject?.image2[0]?.name);
      }
      if (formProject.image3) {
        formData.set('image3', formProject?.image3[0], formProject?.image3[0]?.name);
      }
      if (formProject.image4) {
        formData.set('image4', formProject?.image4[0], formProject?.image4[0]?.name);
      }
      if (formProject.image5) {
        formData.set('image5', formProject?.image5[0], formProject?.image5[0]?.name);
      }
      formData.set('description', formProject.description);

      try {
        const response = await API.post('/project/' + params.id, formData, config);

        const newProject = {
          project: {
            id: response.data.data.id,
            description: response.data.data.description,
            image_1: response.data.data.image_1,
            image_2: response.data.data.image_2,
            image_3: response.data.data.image_3,
            image_4: response.data.data.image_4,
            image_5: response.data.data.image_5,
          }
        }
        props.SetOrders(prevState => prevState.map(order => {
          if (order.id === parseInt(params.id)) {
            return { ...order, ...newProject, };
          }
          return order;
        }));

        setFormProject((formProject) => ({
          ...formProject,
          description: "",
          image1: "",
          image2: "",
          image3: "",
          image4: "",
          image5: "",
        }));
        setImage1("");
        setImage2("");
        setImage3("");
        setImage4("");
        setImage5("");
        try {
          await API.patch('/order/' + params.id, {"status":"complete"});
          const newOrderStatus = {
            status: "complete",
          }
          props.SetOrders(prevOrders => {
            return prevOrders.map(order => {
              if (order.id === params.id) {
                return { ...order, ...newOrderStatus };
              }
              return order;
            });
          });
          props.setMessage();
          props.setShowSuccessUpdateOrderStatusModal();
          navigate("/order-list");
        }
        catch (error) {
          return
        }
      }
      catch (error) {
        return
      }
  });

  return (
    <>
      <Container className="tw-min-h-full">
        <Form onSubmit={(e) => formProjectHandleOnSubmit.mutate(e)} className="md:tw-flex md:tw-justify-between tw-pb-20">
          <div className="animate__animated animate__slideInLeft tw-w-full md:tw-w-6/12 tw-mt-28 md:tw-mt-0">
            {image1 === "" ? (
                <label for="image1" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-mt-28 tw-cursor-pointer">
                  <img src="/images/icon-cloud-upload.png" alt="Project Preview 1" className="tw-object-contain"/>
                  <p><strong className="tw-text-custom-primary">Browse</strong> to choose a file</p>
                </label>
              ) : (
                <label for="image1" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-mt-28 tw-cursor-pointer tw-rounded">
                  <img src={image1} alt="Project Preview 1" className="tw-w-full tw-object-cover"/>
                </label>
              )
            }
            <div className="tw-flex tw-justify-between tw-mt-5">
              {image2 === "" ? (
                <label for="image2" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-2 md:tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer">
                  <img src="/images/icon-add.png" alt="Project Preview 1" className="tw-object-contain"/>
                </label>
              ) : (
                <label for="image2" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer tw-rounded">
                  <img src={image2} alt="Project Preview 1" className="tw-h-full tw-object-cover"/>
                </label>
              )}
              {image3 === "" ? (
                <label for="image3" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-2 md:tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer">
                  <img src="/images/icon-add.png" alt="Project Preview 1" className="tw-object-contain"/>
                </label>
              ) : (
                <label for="image3" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer tw-rounded">
                  <img src={image3} alt="Project Preview 1" className="tw-h-full tw-object-cover"/>
                </label>
              )}
              {image4 === "" ? (
                <label for="image4" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-2 md:tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer">
                  <img src="/images/icon-add.png" alt="Project Preview 1" className="tw-object-contain"/>
                </label>
              ) : (
                <label for="image4" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-mr-4 tw-cursor-pointer tw-rounded">
                  <img src={image4} alt="Project Preview 1" className="tw-h-full tw-object-cover"/>
                </label>
              )}
              {image5 === "" ? (
                <label for="image5" className="tw-border-4 tw-border-custom-secondary-dark hover:tw-border-custom-primary tw-p-2 md:tw-p-5 tw-border-dashed tw-rounded tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-cursor-pointer">
                  <img src="/images/icon-add.png" alt="Project Preview 1" className="tw-object-contain"/>
                </label>
              ) : (
                <label for="image5" className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-w-3/12 tw-cursor-pointer">
                  <img src={image5} alt="Project Preview 1" className="tw-h-full tw-object-cover"/>
                </label>
              )}
            </div>
            <input id="image1" name="image1" onChange={(e) => formProjectHandleOnChangeImage1(e)} type="file" className="tw-hidden"/>
            <input id="image2" name="image2" onChange={(e) => formProjectHandleOnChangeImage2(e)} type="file" className="tw-hidden"/>
            <input id="image3" name="image3" onChange={(e) => formProjectHandleOnChangeImage3(e)} type="file" className="tw-hidden"/>
            <input id="image4" name="image4" onChange={(e) => formProjectHandleOnChangeImage4(e)} type="file" className="tw-hidden"/>
            <input id="image5" name="image5" onChange={(e) => formProjectHandleOnChangeImage5(e)} type="file" className="tw-hidden"/>
          </div>
          <div className="animate__animated animate__slideInRight md:!tw-px-4 !tw-pb-4 tw-flex tw-flex-col tw-items-center tw-mt-10 tw-w-full md:tw-w-5/12 md:tw-mt-28">
            <Form.Group className="!tw-mb-6 tw-w-full" controlId="formDescription">
              <Form.Control as="textarea" onChange={(e) => formProjectHandleOnChange(e)} placeholder="Description" value={formProject.description} name="description" className="!tw-text-lg !tw-p-4 !tw-border-2 !tw-border-solid !tw-border-custom-primary !tw-bg-custom-secondary" rows={8} required/>
            </Form.Group>
            <div className="tw-flex tw-justify-center tw-w-full">
              <Button type="submit" className="!tw-bg-custom-primary hover:!tw-bg-custom-primary-dark !tw-font-bold !tw-border-0 tw-w-full md:tw-w-32">Send Rroject</Button>
            </div>
          </div>
        </Form>
      </Container>
    </>
  );
}