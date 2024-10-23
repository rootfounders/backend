import { CreateProject } from './create_project';
import { CreateProjectForm } from './create_project_form';
import { TipProject } from './tip_project';
import { ProjectComments } from './project_comments';
import { ProjectUpdates } from './project_updates';

(() => {
    console.log('Registrating components');
    customElements.define("create-project", CreateProject);
    customElements.define("create-project-form", CreateProjectForm);
    customElements.define("tip-project", TipProject)
    customElements.define("project-comments", ProjectComments);
    customElements.define("project-updates", ProjectUpdates);
})();

export {
    CreateProject,
};