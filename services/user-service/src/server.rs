use tonic::{transport::Server, Request, Response, Status};

use user::user_server::{User, UserServer};
use user::{UserResponse, UserRequest};

mod user {
    tonic::include_proto!("user"); // The string specified here must match the proto package name
}

#[derive(Debug, Default)]
pub struct MyUser {}

#[tonic::async_trait]
impl User for MyUser {
    async fn get_user(
        &self,
        request: Request<UserRequest>,
    ) -> Result<Response<UserResponse>, Status> {
        println!("Got a request: {:?}", request);

        let response = UserResponse {
            id: "1".to_string(),
            name: "dev day".to_string(),
            username: format!("{}", request.into_inner().username),
            password: "1234".to_string(),
        };

        Ok(Response::new(response)) // Send back our formatted greeting
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let my_user = MyUser::default();

    Server::builder()
        .add_service(UserServer::new(my_user))
        .serve(addr)
        .await?;

    Ok(())
}