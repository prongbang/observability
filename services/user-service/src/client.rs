use user::user_client::UserClient;
use user::UserRequest;

mod user {
    tonic::include_proto!("user");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = UserClient::connect("http://0.0.0.0:50051").await?;

    let request = tonic::Request::new(UserRequest {
        username: "em".into(),
    });

    let response = client.get_user(request).await?;

    println!("RESPONSE={:?}", response);

    Ok(())
}