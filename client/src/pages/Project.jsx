import { Container } from "react-bootstrap";
import { useParams } from "react-router-dom";

export default function Project(props) {
  document.title = "Project | WaysGallery";
  props.setNavbarOn();
  let params = useParams();

  const Order = props.Orders.find(order => order.id === parseInt(params.id));

  return (
    <Container className="tw-min-h-full tw-flex tw-justify-center">
      <div className="tw-flex tw-justify-between tw-my-28">
        <div className="tw-w-7/12">
          {
            Order.project.image_1 === "" ? (
              <img src="/images/image-placeholder" alt="Project Preview 1" className="tw-w-full tw-mb-8 tw-rounded" />
            ) : (
              <img onClick={() => props.handleShowModalImageEnlarger(Order.project.image_1)} src={Order.project.image_1} alt="Project Preview 1" className="animate__animated animate__zoomIn tw-w-full tw-mb-8 tw-rounded tw-cursor-pointer" />
            )
          }
          <div className="tw-flex">
            {
              Order.project.image_2 !== "" && <img onClick={() => props.handleShowModalImageEnlarger(Order.project.image_2)} src={Order.project.image_2} alt="Project Preview 2" className="animate__animated animate__zoomIn tw-w-3/12 tw-h-28 tw-rounded tw-object-cover tw-pr-4 tw-cursor-pointer" />
            }
            {
              Order.project.image_3 !== "" && <img onClick={() => props.handleShowModalImageEnlarger(Order.project.image_3)} src={Order.project.image_3} alt="Project Preview 3" className="animate__animated animate__zoomIn tw-w-3/12 tw-h-28 tw-rounded tw-object-cover tw-pr-4 tw-cursor-pointer" />
            }
            {
              Order.project.image_4 !== "" && <img onClick={() => props.handleShowModalImageEnlarger(Order.project.image_4)} src={Order.project.image_4} alt="Project Preview 4" className="animate__animated animate__zoomIn tw-w-3/12 tw-h-28 tw-rounded tw-object-cover tw-pr-4 tw-cursor-pointer" />
            }
            {
              Order.project.image_5 !== "" && <img onClick={() => props.handleShowModalImageEnlarger(Order.project.image_5)} src={Order.project.image_5} alt="Project Preview 5" className="animate__animated animate__zoomIn tw-w-3/12 tw-h-28 tw-rounded tw-object-cover tw-cursor-pointer" />
            }
          </div>
        </div>
        <p className="animate__animated animate__slideInRight tw-w-4/12">{Order.project.description}</p>
      </div>
    </Container>
  );
}