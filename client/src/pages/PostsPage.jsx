import { Container, Dropdown, Form, InputGroup } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";

function PostsPage(props) {
  document.title = "Posts | WaysGallery";
  const navigate = useNavigate();
  props.setNavbarOn();

  const [postFilter, setPostFilter] = useState("recent");
  const [searchPattern, setSearchPattern] = useState("");

  useEffect(() => {
    const postFilterStorage = localStorage.getItem("WaysGalleryPostFilter");
    if (postFilterStorage === "following" && props.isLogin === true) setPostFilter("following");
    else {
      localStorage.setItem("WaysGalleryPostFilter", "recent");
      setPostFilter("recent")
    };
  }, []);
  const togglePostFilter = () => {
    if (postFilter === "recent") {
      if (props.isLogin) {
        setPostFilter("following");
        localStorage.setItem("WaysGalleryPostFilter", "following");
      }
      else props.showModalLogin();
    }
    else {
      setPostFilter("recent");
      localStorage.setItem("WaysGalleryPostFilter", "recent");
    }
  }
  const Followings = props.Followings.filter(following => following.user.id === props.User.id);

  const PostsData = [...props.Posts];
  PostsData.sort((a, b) => b.id - a.id);
  let Posts = [];
  if (postFilter === "following") {
    for (let post of PostsData) {
      if (Followings.some(following => following.following_id === post.user.id)) Posts.push(post);
    }
  }
  else {
    Posts = PostsData;
  }
  
  Posts = Posts.filter((post) =>
    post.title?.toLowerCase().includes(searchPattern?.toLowerCase())
  );

  // const [PostsState, SetPostsState] = useState(PostsGroup);
  // const searchHandleOnChange = (e) => {
  //   const searchPattern = e.target.value;
  //   const filteredArray = PostsGroup.map((subArray) =>
  //     subArray.filter((item) =>
  //       item.title.toLowerCase().includes(searchPattern.toLowerCase()) || 
  //       item.description.toLowerCase().includes(searchPattern.toLowerCase())
  //     )
  //   );
  //   SetPostsState(filteredArray);
  // };

  const PostsGroup = [[],[],[],[]];
  for (let i = 0; i < Posts.length; i++) {
    if (i % 4 === 0) {
      PostsGroup[0].push(Posts[i]);
    } else if (i % 4 === 1) {
      PostsGroup[1].push(Posts[i]);
    } else if (i % 4 === 2) {
      PostsGroup[2].push(Posts[i]);
    } else {
      PostsGroup[3].push(Posts[i]);
    }
  }

  return (
    <>
      <Container className="tw-min-h-full">
        <div className="tw-flex tw-flex-col md:tw-flex-row tw-justify-between tw-mt-28">
          <Dropdown className="tw-order-2 md:tw-order-1">
            <Dropdown.Toggle className="tw-w-40 !tw-bg-custom-secondary !tw-border-0 hover:!tw-bg-custom-secondary-dark !tw-text-black !tw-font-bold">
              {postFilter === "recent" ? "Recent" : "Following"}
            </Dropdown.Toggle>
            <Dropdown.Menu className="!tw-py-0">
              <Dropdown.Item onClick={togglePostFilter}>{postFilter === "recent" ? "Following" : "Recent"}</Dropdown.Item>
            </Dropdown.Menu>
          </Dropdown>
          <InputGroup className="!tw-mb-4 md:!tw-m-0 md:!tw-w-3/12 tw-order-1 md:tw-order-2">
            <InputGroup.Text className="!tw-bg-custom-secondary !tw-border-0">
              <img src="/images/icon-search.png" alt="search"/>
            </InputGroup.Text>
            <Form.Control value={searchPattern} onChange={(e) => setSearchPattern(e.target.value)} name="pattern" placeholder="Search" aria-label="Search" aria-describedby="search" className="!tw-bg-custom-secondary !tw-border-0"/>
            {/* <Form.Control onChange={(e) => searchHandleOnChange(e)} name="pattern" placeholder="Search" aria-label="Search" aria-describedby="search" className="!tw-bg-custom-secondary !tw-border-0"/> */}
          </InputGroup>
        </div>
        <div>
          <h5 className="tw-font-bold tw-my-10">{postFilter === "recent" ? "Recent Posts" : "Posts You Follow"}</h5>
          <div className="tw-flex tw-flex-wrap">
            {
              PostsData.length > 0 ? (
                PostsGroup.map((postGroup, index) => (
                  postGroup.length > 0 && (
                    <div key={index} className="tw-w-full md:tw-w-3/12">
                      {
                        postGroup.map((post, index) => (
                          <img onClick={() => navigate(`/posts/${post.id}`)} key={index} src={post.image_1} alt={post.title} className="animate__animated animate__zoomIn tw-w-full tw-p-1 tw-object-contain tw-border-md hover:tw-opacity-75 tw-cursor-pointer" />
                        ))
                      }
                    </div>
                  )
                ))
              ) : (
                <p className="tw-opacity-50">There are no posts to display.</p>
              )
            }
          </div>
        </div>
      </Container>
    </>
  )
}

export default PostsPage;