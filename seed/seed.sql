-- seeds.sql
DO $$
DECLARE
    user1_id UUID;
    user2_id UUID;
    project1_id UUID;
    project2_id UUID;
BEGIN
    -- Users
    INSERT INTO users (id, github_id, name, email) VALUES 
    (gen_random_uuid(), 'user1', 'Test User 1', 'user1@example.com') RETURNING id INTO user1_id;
    INSERT INTO users (id, github_id, name, email) VALUES 
    (gen_random_uuid(), 'user2', 'Test User 2', 'user2@example.com') RETURNING id INTO user2_id;

    -- Projects
    INSERT INTO projects (id, github_repo_id, name, description) VALUES 
    (gen_random_uuid(), 'repo1', 'Test Project 1', 'This is a test project 1') RETURNING id INTO project1_id;
    INSERT INTO projects (id, github_repo_id, name, description) VALUES 
    (gen_random_uuid(), 'repo2', 'Test Project 2', 'This is a test project 2') RETURNING id INTO project2_id;

    -- Project Contributors
    INSERT INTO project_contributors (project_id, user_id, role, can_manage_job_posting) VALUES 
    (project1_id, user1_id, 'owner', true),
    (project2_id, user2_id, 'owner', true);

END $$;