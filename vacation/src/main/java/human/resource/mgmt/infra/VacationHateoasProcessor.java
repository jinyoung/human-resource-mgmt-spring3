package human.resource.mgmt.infra;

import human.resource.mgmt.domain.*;
import org.springframework.hateoas.EntityModel;
import org.springframework.hateoas.Link;
import org.springframework.hateoas.server.RepresentationModelProcessor;
import org.springframework.stereotype.Component;

@Component
public class VacationHateoasProcessor
    implements RepresentationModelProcessor<EntityModel<Vacation>> {

    @Override
    public EntityModel<Vacation> process(EntityModel<Vacation> model) {
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/cancel")
                .withRel("cancel")
        );
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/approve")
                .withRel("approve")
        );
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/confirmused")
                .withRel("confirmused")
        );
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/update")
                .withRel("update")
        );

        return model;
    }
}
